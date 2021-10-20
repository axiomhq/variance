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
	assert.EqualValues(t, 0.3, s.VarianceSample())
	assert.EqualValues(t, 0.5477225575051661, s.StandardDeviationSample())

	s.Clear()
	s.Add(3)
	s.Add(6)
	s.Add(4)
	s.Add(5)
	s.Add(5)
	s.Add(4)

	assert.EqualValues(t, 4.5, s.Mean())
	assert.EqualValues(t, 1.1, s.VarianceSample())
	assert.EqualValues(t, 1.0488088481701516, s.StandardDeviationSample())
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
	assert.EqualValues(t, sTotal.VarianceSample(), s1.VarianceSample())

	s1.Add(5555)
	sTotal.Add(5555)

	assert.EqualValues(t, sTotal.Mean(), s1.Mean())
	assert.EqualValues(t, sTotal.VarianceSample(), s1.VarianceSample())
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
	assert.InDelta(t, 1.0, s.VarianceSample(), 0.001)
	assert.InDelta(t, 1.0, s.StandardDeviationSample(), 0.001)
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
