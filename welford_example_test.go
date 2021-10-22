package variance_test

import (
	"fmt"

	"github.com/axiomhq/variance"
)

func Example() {
	stats1 := variance.New()

	stats1.Add(1)
	stats1.Add(1)
	stats1.Add(1)
	stats1.Add(0)
	stats1.Add(0)
	stats1.Add(0)

	fmt.Println(
		stats1.Mean(),
		stats1.Variance(),
		stats1.StandardDeviation(),
		stats1.VariancePopulation(),
		stats1.StandardDeviationPopulation(),
		stats1.NumDataValues(),
	)

	stats2 := variance.New()
	stats2.Add(3)

	// Merge the values of stats2 into stats1.
	stats1.Merge(stats2)

	// Reset the values in stats2.
	stats2.Clear()

	// Output: 0.5 0.3 0.5477225575051661 0.25 0.5 6
}
