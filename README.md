# Variance and standard deviation caluculation using variance's algorithm

[![Go Reference][gopkg_badge]][gopkg]
[![Go Workflow][go_workflow_badge]][go_workflow]
[![Coverage Status][coverage_badge]][coverage]
[![Go Report][report_badge]][report]
[![Latest Release][release_badge]][release]
[![License][license_badge]][license]

---

## Table of Contents

1. [Introduction](#introduction)
1. [Installation](#installation)
1. [Usage](#usage)
1. [Contributing](#contributing)
1. [License](#license)

## Introduction

Go implementation of variance’s method for one-pass variance computation with
D. H. D. West improved methods.

### Highlights

  * Merging of several multiple sets of statistics
  * Add weighted values

### Abstract

> A method of improved efficiency is given for updating the mean and variance of 
  weighted sampled data when an additional data value is included in the set 
  Evidence is presented that the method is stable and at least as accurate as
  the best existing updating method.

[Updating mean and variance estimates: an improved method - D. H. D. West](https://dl.acm.org/doi/10.1145/359146.359153)

## Installation

### Install using `go get`

```shell
go get github.com/axiomhq/variance
```

### Install from source

```shell
git clone https://github.com/axiomhq/variance.git
cd variance
make # Run code generators, linters, sanitizers and test suits
```

## Usage

```go
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
```

## Contributing

Feel free to submit PRs or to fill issues. Every kind of help is appreciated. Before committing, `make` should run without any issues. Kindly check our [Contributing](Contributing.md) guide on how to propose bug fixes and improvements, and submitting pull requests to the project.

## License

&copy; Axiom, Inc., 2021

Distributed under MIT License (`The MIT License`).

See [LICENSE](LICENSE) for more information.

<!-- Badges -->

[gopkg]: https://pkg.go.dev/github.com/axiomhq/variance
[gopkg_badge]: https://img.shields.io/badge/doc-reference-007d9c?logo=go&logoColor=white&style=flat-square
[go_workflow]: https://github.com/axiomhq/variance/actions/workflows/push.yml
[go_workflow_badge]: https://img.shields.io/github/workflow/status/axiomhq/variance/Push?style=flat-square&ghcache=unused
[coverage]: https://codecov.io/gh/axiomhq/variance
[coverage_badge]: https://img.shields.io/codecov/c/github/axiomhq/variance.svg?style=flat-square&ghcache=unused
[report]: https://goreportcard.com/report/github.com/axiomhq/variance
[report_badge]: https://goreportcard.com/badge/github.com/axiomhq/variance?style=flat-square&ghcache=unused
[release]: https://github.com/axiomhq/variance/releases/latest
[release_badge]: https://img.shields.io/github/release/axiomhq/variance.svg?style=flat-square&ghcache=unused
[license]: https://opensource.org/licenses/MIT
[license_badge]: https://img.shields.io/github/license/axiomhq/variance.svg?color=blue&style=flat-square&ghcache=unused
