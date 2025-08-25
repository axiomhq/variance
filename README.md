# variance [![Go Reference][gopkg_badge]][gopkg] [![Workflow][workflow_badge]][workflow] [![Latest Release][release_badge]][release] [![License][license_badge]][license]

Go implementation of Welford's method for one-pass variance computation with D. H. D. West improved methods.

## Install

```shell
go get github.com/axiomhq/variance
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/axiomhq/variance"
)

func main() {
    stats := variance.New()

    stats.Add(1)
    stats.Add(1)
    stats.Add(1)
    stats.Add(0)
    stats.Add(0)
    stats.Add(0)

    fmt.Println(
        stats.Mean(),
        stats.Variance(),
        stats.StandardDeviation(),
        stats.VariancePopulation(),
        stats.StandardDeviationPopulation(),
        stats.NumDataValues(),
    )
}
```

For more examples, check out the [example](welford_example_test.go) or
[run it on pkg.go.dev](https://pkg.go.dev/github.com/axiomhq/variance#example-package).

## Features

- One-pass variance computation
- Merging of multiple sets of statistics
- Support for weighted values
- Stable and accurate computation using D. H. D. West's improved method

## Reference

> A method of improved efficiency is given for updating the mean and variance of
> weighted sampled data when an additional data value is included in the set.
> Evidence is presented that the method is stable and at least as accurate as
> the best existing updating method.
>
> -- <cite>[Updating mean and variance estimates: an improved method - D. H. D. West](https://dl.acm.org/doi/10.1145/359146.359153)</cite>

## License

[MIT](LICENSE)

<!-- Badges -->

[gopkg]: https://pkg.go.dev/github.com/axiomhq/variance
[gopkg_badge]: https://img.shields.io/badge/doc-reference-007d9c?logo=go&logoColor=white
[workflow]: https://github.com/axiomhq/variance/actions/workflows/push.yaml
[workflow_badge]: https://img.shields.io/github/actions/workflow/status/axiomhq/variance/push.yaml?branch=main&ghcache=unused
[release]: https://github.com/axiomhq/variance/releases/latest
[release_badge]: https://img.shields.io/github/release/axiomhq/variance.svg?ghcache=unused
[license]: https://opensource.org/licenses/MIT
[license_badge]: https://img.shields.io/github/license/axiomhq/variance.svg?color=blue&ghcache=unused
