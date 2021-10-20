# welford
Go implementation Welford’s method for one-pass variance computation with D. H. D. West improved method

## Abstract

> A method of improved efficiency is given for updating the mean and variance of weighted sampled data when an additional data value is included in the set. Evidence is presented 
> that the method is stable and at least as accurate as the best existing updating method.

## Example Usage

```go

sk1 := welford.New()
sk1.Add(1)
sk1.Add(1)
sk1.Add(1)
sk1.Add(0)
sk1.Add(0)
sk1.Add(0)

mean := s.Mean() // ==> 0.5
variance := s.Variance() // ==> 0.3
stdev := s.StandardDeviation() // ==> 0.5477225575051661
variancp := s.VariancePpopulation()	// ==> 0.25
stdevp := s.StandardDeviationPopulation() // ==> 0.5
n := s.NumDataValues()) // 6

sk2 := welford.New()
sk2.Add(3)

sk1.Merge(sk2) // merge sk1 into sk2

sk2.Clear() // resets the state os sk2
```