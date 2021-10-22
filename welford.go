package variance

import (
	"encoding/binary"
	"io"
	"math"
)

var (
	_ io.ReaderFrom = (*Stats)(nil)
	_ io.WriterTo   = (*Stats)(nil)
)

// Stats of welford's algorithm for computing the mean and variance.
type Stats struct {
	n    uint
	mean float64
	sum  float64
	sum2 float64
	s    float64
}

// New returns new Welford's algorithm stats. The stats are initialized to their
// respective zero values.
func New() *Stats {
	return new(Stats)
}

// Add a new value to the stats.
func (sts *Stats) Add(x float64) {
	sts.AddWeighted(x, 1)
}

// AddWeighted adds a new weighted value to the stats.
func (sts *Stats) AddWeighted(val, weight float64) {
	sts.n++
	sts.sum += weight
	sts.sum2 += weight * weight
	meanOld := sts.mean
	sts.mean = meanOld + (weight/sts.sum)*(val-meanOld)
	sts.s = sts.s + weight*(val-meanOld)*(val-sts.mean)
}

// Clear the stats to its initial state.
func (sts *Stats) Clear() {
	sts.n = 0
	sts.s = 0
	sts.sum = 0
	sts.sum2 = 0
	sts.mean = 0
}

// Mean returns the mean of the data.
func (sts *Stats) Mean() float64 {
	return sts.mean
}

// VariancePopulation returns the variance of the data, assuming the data added
// was not sampled.
func (sts *Stats) VariancePopulation() float64 {
	return sts.s / sts.sum
}

// Variance returns the variance of the data, assuming the data added was
// sampled.
func (sts *Stats) Variance() float64 {
	return sts.s / (sts.sum - 1)
}

// StandardDeviationPopulation returns the standard deviation of the data, which
// is the square root of the variance of the sampled data.
func (sts *Stats) StandardDeviationPopulation() float64 {
	return math.Sqrt(sts.VariancePopulation())
}

// StandardDeviation returns the standard deviation of the data, which is the
// square root of the variance.
func (sts *Stats) StandardDeviation() float64 {
	return math.Sqrt(sts.Variance())
}

// NumDataValues returns the number of data values in the stats.
func (sts *Stats) NumDataValues() uint {
	return sts.n
}

// Clone returns a copy of the stats.
func (sts *Stats) Clone() *Stats {
	return &Stats{
		n:    sts.n,
		mean: sts.mean,
		sum:  sts.sum,
		sum2: sts.sum2,
		s:    sts.s,
	}
}

// Merge other stats into the stats.
func (sts *Stats) Merge(other *Stats) {
	sts.n += other.n
	sts.sum += other.sum
	sts.sum2 += other.sum2
	meanOld := sts.mean
	sts.mean = meanOld + (other.sum/sts.sum)*(other.mean-meanOld)
	sts.s = sts.s + other.s + other.sum*(other.mean-meanOld)*(other.mean-sts.mean)
}

// ReadFrom reads the stats from `r`.
func (sts *Stats) ReadFrom(r io.Reader) (int64, error) {
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

	sts.n = uint(n)
	sts.mean = mean
	sts.sum = sum
	sts.sum2 = sum2
	sts.s = s

	return 40, nil
}

// WriteTo writes the stats to `w`.
func (sts *Stats) WriteTo(w io.Writer) (int64, error) {
	var err error

	if err = binary.Write(w, binary.BigEndian, uint64(sts.n)); err != nil {
		return 0, err
	}
	if err = binary.Write(w, binary.BigEndian, sts.mean); err != nil {
		return 8, err
	}
	if err = binary.Write(w, binary.BigEndian, sts.sum); err != nil {
		return 16, err
	}
	if err = binary.Write(w, binary.BigEndian, sts.sum2); err != nil {
		return 24, err
	}
	if err = binary.Write(w, binary.BigEndian, sts.s); err != nil {
		return 32, err
	}

	return 40, nil
}
