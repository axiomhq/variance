package welford

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatsSimple(t *testing.T) {
	suite := assert.New(t)

	stats := New()
	stats.Add(1)
	stats.Add(1)
	stats.Add(1)
	stats.Add(0)
	stats.Add(0)
	stats.Add(0)

	suite.EqualValues(0.5, stats.Mean())
	suite.EqualValues(0.3, stats.Variance())
	suite.EqualValues(0.5477225575051661, stats.StandardDeviation())
	suite.EqualValues(0.25, stats.VariancePopulation())
	suite.EqualValues(0.5, stats.StandardDeviationPopulation())
	suite.EqualValues(6, stats.NumDataValues())

	stats.Clear()
	stats.Add(3)
	stats.Add(6)
	stats.Add(4)
	stats.Add(5)
	stats.Add(5)
	stats.Add(4)

	suite.EqualValues(4.5, stats.Mean())
	suite.EqualValues(1.1, stats.Variance())
	suite.EqualValues(1.0488088481701516, stats.StandardDeviation())
	suite.EqualValues(0.9166666666666666, stats.VariancePopulation())
	suite.EqualValues(0.9574271077563381, stats.StandardDeviationPopulation())
	suite.EqualValues(6, stats.NumDataValues())

	stats.Clear()
	stats.Add(13)
	stats.Add(17)
	stats.Add(18)
	stats.Add(27)
	stats.Add(28)

	suite.EqualValues(20.6, stats.Mean())
	suite.EqualValues(43.3, stats.Variance())
	suite.EqualValues(6.58027355054484, stats.StandardDeviation())
	suite.EqualValues(34.64, stats.VariancePopulation())
	suite.EqualValues(5.885575587824865, stats.StandardDeviationPopulation())
	suite.EqualValues(5, stats.NumDataValues())
}

func TestStats_Merge(t *testing.T) {
	suite := assert.New(t)

	var (
		statsTotal = New()
		stats1     = New()
		stats2     = New()
	)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			stats1.Add(float64(i))
		} else {
			stats2.Add(float64(i))
		}
		statsTotal.Add(float64(i))
	}

	stats1.Merge(stats2)
	suite.EqualValues(statsTotal.Mean(), stats1.Mean())
	suite.EqualValues(statsTotal.Variance(), stats1.Variance())

	stats1.Add(5555)
	statsTotal.Add(5555)

	suite.EqualValues(statsTotal.Mean(), stats1.Mean())
	suite.EqualValues(statsTotal.Variance(), stats1.Variance())
}

func TestStats_NumDataValues(t *testing.T) {
	stats := New()
	num := rand.Intn(100) //nolint:gosec // Fine for testing

	for i := 0; i < num; i++ {
		stats.Add(1)
	}

	assert.EqualValues(t, num, stats.NumDataValues())
}

func TestStatsWithRandomFloats(t *testing.T) {
	suite := assert.New(t)

	stats := New()

	for i := 0; i < 1000000; i++ {
		stats.Add(rand.NormFloat64())
	}

	// Allow .1% of error (.1% arbitrarily chosen)
	suite.InDelta(0.0, stats.Mean(), 0.001)
	suite.InDelta(1.0, stats.Variance(), 0.001)
	suite.InDelta(1.0, stats.StandardDeviation(), 0.001)
}

func TestStatsWriteToReadFrom(t *testing.T) {
	suite := assert.New(t)

	exp := New()
	exp.Add(1)
	exp.Add(2)
	exp.Add(3)

	var buf bytes.Buffer

	n, err := exp.WriteTo(&buf)
	suite.NoError(err)
	suite.EqualValues(40, n)

	got := New()
	n, err = got.ReadFrom(&buf)
	suite.NoError(err)
	suite.EqualValues(40, n)

	suite.EqualValues(exp, got)
}
