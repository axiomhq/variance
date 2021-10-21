package welford

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWelfordSimple(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add(1)
	s.Add(1)
	s.Add(0)
	s.Add(0)
	s.Add(0)

	assert.EqualValues(t, 0.5, s.Mean())
	assert.EqualValues(t, 0.3, s.Variance())
	assert.EqualValues(t, 0.5477225575051661, s.StandardDeviation())
	assert.EqualValues(t, 0.25, s.VariancePopulation())
	assert.EqualValues(t, 0.5, s.StandardDeviationPopulation())
	assert.EqualValues(t, 6, s.NumDataValues())

	s.Clear()
	s.Add(3)
	s.Add(6)
	s.Add(4)
	s.Add(5)
	s.Add(5)
	s.Add(4)

	assert.EqualValues(t, 4.5, s.Mean())
	assert.EqualValues(t, 1.1, s.Variance())
	assert.EqualValues(t, 1.0488088481701516, s.StandardDeviation())
	assert.EqualValues(t, 0.9166666666666666, s.VariancePopulation())
	assert.EqualValues(t, 0.9574271077563381, s.StandardDeviationPopulation())
	assert.EqualValues(t, 6, s.NumDataValues())

	s.Clear()
	s.Add(13)
	s.Add(17)
	s.Add(18)
	s.Add(27)
	s.Add(28)

	assert.EqualValues(t, 20.6, s.Mean())
	assert.EqualValues(t, 43.3, s.Variance())
	assert.EqualValues(t, 6.58027355054484, s.StandardDeviation())
	assert.EqualValues(t, 34.64, s.VariancePopulation())
	assert.EqualValues(t, 5.885575587824865, s.StandardDeviationPopulation())
	assert.EqualValues(t, 5, s.NumDataValues())
}

func TestMomentsSimpleMerge(t *testing.T) {
	var (
		sTotal = New()
		s1     = New()
		s2     = New()
	)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			s1.Add(float64(i))
		} else {
			s2.Add(float64(i))
		}
		sTotal.Add(float64(i))
	}

	s1.Merge(*s2)
	assert.EqualValues(t, sTotal.Mean(), s1.Mean())
	assert.EqualValues(t, sTotal.Variance(), s1.Variance())

	s1.Add(5555)
	sTotal.Add(5555)

	assert.EqualValues(t, sTotal.Mean(), s1.Mean())
	assert.EqualValues(t, sTotal.Variance(), s1.Variance())
}

func TestMomentsNumDataValues(t *testing.T) {
	s := New()
	num := rand.Intn(100)
	for i := 0; i < num; i++ {
		s.Add(1)
	}
	assert.EqualValues(t, num, s.NumDataValues())
}

func TestMomentsRandomFloats(t *testing.T) {
	s := New()
	for i := 0; i < 1000000; i++ {
		s.Add(rand.NormFloat64())
	}
	// Allow .1% of error (.1% arbitrarily chosen)
	assert.InDelta(t, 0.0, s.Mean(), 0.001)
	assert.InDelta(t, 1.0, s.Variance(), 0.001)
	assert.InDelta(t, 1.0, s.StandardDeviation(), 0.001)
}

func TestMomentsEncodeDecode(t *testing.T) {
	exp := New()
	exp.Add(1)
	exp.Add(2)
	exp.Add(3)

	p := bytes.NewBuffer(nil)

	n, err := exp.WriteTo(p)
	assert.NoError(t, err)
	assert.EqualValues(t, 40, n)

	got := New()
	n, err = got.ReadFrom(p)

	assert.NoError(t, err)
	assert.EqualValues(t, 40, n)
	assert.EqualValues(t, exp, got)
}
