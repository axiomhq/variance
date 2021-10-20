# Welford - Online/Rolling method of calculating variance and standard deviation
Go implementation Welford’s method for one-pass variance computation with D. H. D. West improved methods.

Highlights:
* Merging of several multiple sets of statistics
* Add weighted values

## Abstract

> A method of improved efficiency is given for updating the mean and variance of weighted sampled data when an additional data value is included in the set. Evidence is presented 
> that the method is stable and at least as accurate as the best existing updating method.

- [Updating mean and variance estimates: an improved method - D. H. D. West](https://dl.acm.org/doi/10.1145/359146.359153)

## Example Usage

```go

sts1 := welford.New()
sts1.Add(1)
sts1.Add(1)
sts1.Add(1)
sts1.Add(0)
sts1.Add(0)
sts1.Add(0)

mean := s.Mean() // ==> 0.5
variance := s.Variance() // ==> 0.3
stdev := s.StandardDeviation() // ==> 0.5477225575051661
variancp := s.VariancePpopulation()	// ==> 0.25
stdevp := s.StandardDeviationPopulation() // ==> 0.5
n := s.NumDataValues()) // 6

sts2 := welford.New()
sts2.Add(3)

sts1.Merge(sts2) // merge sts1 into sts2

sts2.Clear() // resets the state os sts2
```
