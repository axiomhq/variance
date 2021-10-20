package welford

import (
	"encoding/binary"
	"io"
	"math"
)

// Welford's algorithm for computing the mean and variance.
type Sketch struct {
	n    uint
	mean float64
	sum  float64
	sum2 float64
	s    float64
}

// NewSketch returns a new Welford's algorithm sketch.
func New() *Sketch {
	return &Sketch{}
}

// Add a new value to the sketch.
func (sk *Sketch) Add(x float64) {
	sk.AddWeighted(x, 1)
}

// Add a new weighted value to the sketch.
func (sk *Sketch) AddWeighted(val, weight float64) {
	sk.n++
	sk.sum += weight
	sk.sum2 += weight * weight
	meanOld := sk.mean
	sk.mean = meanOld + (weight/sk.sum)*(val-meanOld)
	sk.s = sk.s + weight*(val-meanOld)*(val-sk.mean)
}

// Clear resets the sketch to its initial state.
func (sk *Sketch) Clear() {
	sk.n = 0
	sk.s = 0
	sk.sum = 0
	sk.sum2 = 0
	sk.mean = 0
}

// Mean returns the mean of the data.
func (sk *Sketch) Mean() float64 {
	return sk.mean
}

// VariancePopulation returns the variance of the data, assuming the data added was not sampled.
func (sk *Sketch) VariancePopulation() float64 {
	return sk.s / sk.sum
}

// Variance returns the variance of the data, assuming the data added was sampled.
func (sk *Sketch) Variance() float64 {
	return sk.s / (sk.sum - 1)
}

// Standard deviation is the square root of the variance.
func (sk *Sketch) StandardDeviationPopulation() float64 {
	return math.Sqrt(sk.VariancePopulation())
}

// Standard deviation is the square root of the variance.
func (sk *Sketch) StandardDeviation() float64 {
	return math.Sqrt(sk.Variance())
}

// NumDataValues returns the number of data values in the sketch.
func (sk *Sketch) NumDataValues() uint {
	return sk.n
}

// Clone returns a copy of the sketch.
func (sk *Sketch) Clone() *Sketch {
	return &Sketch{
		n:    sk.n,
		mean: sk.mean,
		sum:  sk.sum,
		sum2: sk.sum2,
		s:    sk.s,
	}
}

// Merge `other` sketch into the receiver.
func (sk *Sketch) Merge(other Sketch) {
	sk.n += other.n
	sk.sum += other.sum
	sk.sum2 += other.sum2
	meanOld := sk.mean
	sk.mean = meanOld + (other.sum/sk.sum)*(other.mean-meanOld)
	sk.s = sk.s + other.s + other.sum*(other.mean-meanOld)*(other.mean-sk.mean)
}

// WriteTo writes the sketch to `w`.
func (sk *Sketch) WriteTo(w io.Writer) (int64, error) {
	if err := binary.Write(w, binary.BigEndian, uint64(sk.n)); err != nil {
		return 0, err
	}
	if err := binary.Write(w, binary.BigEndian, sk.mean); err != nil {
		return 8, err
	}
	if err := binary.Write(w, binary.BigEndian, sk.sum); err != nil {
		return 16, err
	}
	if err := binary.Write(w, binary.BigEndian, sk.sum2); err != nil {
		return 24, err
	}
	if err := binary.Write(w, binary.BigEndian, sk.s); err != nil {
		return 32, err
	}
	return 40, nil
}

// ReadFrom reads the sketch from `r`.
func (sk *Sketch) ReadFrom(r io.Reader) (int64, error) {
	var (
		n                  uint64
		mean, sum, sum2, s float64
		err                error
	)
	if err = binary.Read(r, binary.BigEndian, &n); err != nil {
		return 0, err
	}
	if err = binary.Read(r, binary.BigEndian, &mean); err != nil {
		return 8, err
	}
	if err = binary.Read(r, binary.BigEndian, &sum); err != nil {
		return 16, err
	}
	if err = binary.Read(r, binary.BigEndian, &sum2); err != nil {
		return 24, err
	}
	if err = binary.Read(r, binary.BigEndian, &s); err != nil {
		return 32, err
	}
	sk.n = uint(n)
	sk.mean = mean
	sk.sum = sum
	sk.sum2 = sum2
	sk.s = s
	return 40, nil
}
