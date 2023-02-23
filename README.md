![variance: Variance and standard deviation calculation using variance's algorithm](.github/images/banner-dark.svg#gh-dark-mode-only)
![variance: Variance and standard deviation calculation using variance's algorithm](.github/images/banner-light.svg#gh-light-mode-only)

<div align="center">

[![Go Reference][gopkg_badge]][gopkg]
[![Workflow][workflow_badge]][workflow]
[![Latest Release][release_badge]][release]
[![License][license_badge]][license]

</div>

[Axiom](https://axiom.co) unlocks observability at any scale.

- **Ingest with ease, store without limits:** Axiom's next-generation datastore
  enables ingesting petabytes of data with ultimate efficiency. Ship logs from
  Kubernetes, AWS, Azure, Google Cloud, DigitalOcean, Nomad, and others.
- **Query everything, all the time:** Whether DevOps, SecOps, or EverythingOps,
  query all your data no matter its age. No provisioning, no moving data from
  cold/archive to "hot", and no worrying about slow queries. All your data, all.
  the. time.
- **Powerful dashboards, for continuous observability:** Build dashboards to
  collect related queries and present information that's quick and easy to
  digest for you and your team. Dashboards can be kept private or shared with
  others, and are the perfect way to bring together data from different sources.

For more information check out the
[official documentation](https://axiom.co/docs) and our
[community Discord](https://axiom.co/discord).

## Introduction

Go implementation of variance's method for one-pass variance computation with
D. H. D. West improved methods:

> A method of improved efficiency is given for updating the mean and variance of
> weighted sampled data when an additional data value is included in the set.
> Evidence is presented that the method is stable and at least as accurate as
> the best existing updating method.
>
> -- <cite>[Updating mean and variance estimates: an improved method - D. H. D. West](https://dl.acm.org/doi/10.1145/359146.359153)</cite>

It features merging of several multiple sets of statistics and adding weighted
values.

## Quickstart

Install using `go get`:

```shell
go get github.com/axiomhq/variance
```

Import the package:

```go
import "github.com/axiomhq/variance"
```

Use the package:

```go
package main

import (
 "fmt"

 "github.com/axiomhq/variance"
)

func main() {
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
}
```

Checkout the [example](welford_example_test.go) or
[run it on pkg.go.dev](https://pkg.go.dev/github.com/axiomhq/variance#example-package).

## License

Distributed under the [MIT License](LICENSE).

<!-- Badges -->

[gopkg]: https://pkg.go.dev/github.com/axiomhq/variance
[gopkg_badge]: https://img.shields.io/badge/doc-reference-007d9c?logo=go&logoColor=white
[workflow]: https://github.com/axiomhq/variance/actions/workflows/push.yaml
[workflow_badge]: https://img.shields.io/github/actions/workflow/status/axiomhq/variance/push.yaml?branch=main&ghcache=unused
[release]: https://github.com/axiomhq/variance/releases/latest
[release_badge]: https://img.shields.io/github/release/axiomhq/variance.svg?ghcache=unused
[license]: https://opensource.org/licenses/MIT
[license_badge]: https://img.shields.io/github/license/axiomhq/variance.svg?color=blue&ghcache=unused
