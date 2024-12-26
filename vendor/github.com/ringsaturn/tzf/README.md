# TZF: a fast timezone finder for Go. [![Go Reference](https://pkg.go.dev/badge/github.com/ringsaturn/tzf.svg)](https://pkg.go.dev/github.com/ringsaturn/tzf) [![codecov](https://codecov.io/gh/ringsaturn/tzf/branch/main/graph/badge.svg?token=9KIU85IERM)](https://codecov.io/gh/ringsaturn/tzf)

![](https://github.com/ringsaturn/tzf/blob/gh-pages/docs/tzf-social-media.png?raw=true)

TZF is a fast timezone finder package designed for Go. It allows you to quickly
find the timezone for a given latitude and longitude, making it ideal for geo
queries and services such as weather forecast APIs. With optimized performance
and two different data options, TZF is a powerful tool for any Go developer's
toolkit.

---

> [!NOTE]
>
> Here are some language or server which built with tzf or it's other language
> bindings:

| Language or Sever | Link                                                                    | Note              |
| ----------------- | ----------------------------------------------------------------------- | ----------------- |
| Go                | [`ringsaturn/tzf`](https://github.com/ringsaturn/tzf)                   |                   |
| Ruby              | [`HarlemSquirrel/tzf-rb`](https://github.com/HarlemSquirrel/tzf-rb)     |                   |
| Rust              | [`ringsaturn/tzf-rs`](https://github.com/ringsaturn/tzf-rs)             |                   |
| Python            | [`ringsaturn/tzfpy`](https://github.com/ringsaturn/tzfpy)               |                   |
| HTTP API          | [`ringsaturn/tzf-server`](https://github.com/ringsaturn/tzf-server)     | build with tzf    |
| HTTP API          | [`racemap/rust-tz-service`](https://github.com/racemap/rust-tz-service) | build with tzf-rs |
| Redis Server      | [`ringsaturn/tzf-server`](https://github.com/ringsaturn/tzf-server)     | build with tzf    |
| Redis Server      | [`ringsaturn/redizone`](https://github.com/ringsaturn/redizone)         | build with tzf-rs |

## Quick Start

To start using TZF in your Go project, you first need to install the package:

```bash
go get github.com/ringsaturn/tzf
```

Then, you can use the following code to locate:

```go
// Use about 150MB memory for init, and 60MB after GC.
package main

import (
	"fmt"

	"github.com/ringsaturn/tzf"
)

func main() {
	finder, err := tzf.NewDefaultFinder()
	if err != nil {
		panic(err)
	}
	fmt.Println(finder.GetTimezoneName(116.6386, 40.0786))
}
```

If you require a query result that is 100% accurate, use the following to
locate:

```go
// Use about 900MB memory for init, and 660MB after GC.
package main

import (
	"fmt"

	"github.com/ringsaturn/tzf"
	tzfrel "github.com/ringsaturn/tzf-rel"
	"github.com/ringsaturn/tzf/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	input := &pb.Timezones{}

	// Full data, about 83.5MB
	dataFile := tzfrel.FullData

	if err := proto.Unmarshal(dataFile, input); err != nil {
		panic(err)
	}
	finder, _ := tzf.NewFinderFromPB(input)
	fmt.Println(finder.GetTimezoneName(116.6386, 40.0786))
}
```

### Best Practice

It's expensive to init tzf's Finder/FuzzyFinder/DefaultFinder, please consider
reuse it or as a global var. Below is a global var example:

```go
package main

import (
	"fmt"

	"github.com/ringsaturn/tzf"
)

var f tzf.F

func init() {
	var err error
	f, err = tzf.NewDefaultFinder()
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(f.GetTimezoneName(116.3883, 39.9289))
	fmt.Println(f.GetTimezoneName(-73.935242, 40.730610))
}
```

## CLI Tool

In addition to using TZF as a library in your Go projects, you can also use the
tzf command-line interface (CLI) tool to quickly get the timezone name for a set
of coordinates. To use the CLI tool, you first need to install it using the
following command:

```bash
go install github.com/ringsaturn/tzf/cmd/tzf@latest
```

Once installed, you can use the tzf command followed by the latitude and
longitude values to get the timezone name:

```bash
tzf -lng 116.3883 -lat 39.9289
```

## Data

You can download the original data from
<https://github.com/evansiroky/timezone-boundary-builder>.

The preprocessed protobuf data can be obtained from
<https://github.com/ringsaturn/tzf-rel>, which has Go's `embedded` support.
These files are Protocol Buffers messages for more efficient binary
distribution, similar to Python wheels. You can view the
[`pb/tzinfo.proto file`](./pb/tzinfo.proto) or its
[HTML format documentation][pb_html] for information about the internal format.

The data pipeline for tzf can be illustrated as follows:

```mermaid
graph TD
    Raw[GeoJSON from evansiroky/timezone-boundary-builder]
    Full[Full: Probuf based data]
    Lite[Lite: smaller of Full data]
    Compressed[Compressed: Lite compressed via Polyline]
    Preindex[Tile based data]

    Finder[Finder: Polygon Based Finder]
    FuzzyFinder[FuzzyFinder: Tile based Finder]
    DefaultFinder[DefaultFinder: combine FuzzyFinder and Compressed Finder]

    Raw --> |cmd/geojson2tzpb|Full
    Full --> |cmd/reducetzpb|Lite
    Lite --> |cmd/compresstzpb|Compressed
    Lite --> |cmd/preindextzpb|Preindex

    Full --> |tzf.NewFinderFromPB|Finder
    Lite --> |tzf.NewFinderFromPB|Finder
    Compressed --> |tzf.NewFinderFromCompressed|Finder --> |tzf.NewDefaultFinder|DefaultFinder
    Preindex --> |tzf.NewFuzzyFinderFromPB|FuzzyFinder --> |tzf.NewDefaultFinder|DefaultFinder
```

The [complete dataset (~80MB)][full-link] can be used anywhere, but requires
higher memory usage.

The [lightweight dataset (~10MB)][lite-link] may not function optimally in some
border areas.

You can observe points with different outcomes on this [page][points_not_equal].

If a slightly longer initialization time is tolerable, the
[compressed dataset (~5MB)][compressd-link] derived from the lightweight dataset
will be **more suitable for binary distribution.**

The [pre-indexed dataset (~1.78MB)][preindex-link] consists of multiple tiles.
It is used within the `DefaultFinder`, which is built on `FuzzyFinder`, to
reduce execution times of the raycasting algorithm.

[pb_html]: https://ringsaturn.github.io/tzf/pb.html
[full-link]: https://github.com/ringsaturn/tzf-rel/blob/main/combined-with-oceans.pb
[lite-link]: https://github.com/ringsaturn/tzf-rel/blob/main/combined-with-oceans.reduce.pb
[preindex-link]: https://github.com/ringsaturn/tzf-rel/blob/main/combined-with-oceans.reduce.preindex.pb
[compressd-link]: https://github.com/ringsaturn/tzf-rel/blob/main/combined-with-oceans.reduce.compress.pb
[points_not_equal]: https://geojson.io/#id=gist:ringsaturn/2d958e7f0a279a7411c04907f255955a

I have written an article about the history of tzf, its Rust port, and its Rust
port's Python binding; you can view it
[here](https://blog.ringsaturn.me/en/posts/2023-01-31-history-of-tzf/).

## Performance

The tzf package is intended for high-performance geospatial query services, such
as weather forecasting APIs. Most queries can be returned within a very short
time, averaging around 2000 nanoseconds.

Here is what has been done to improve performance:

1. Using pre-indexing to handle most queries takes approximately 1000
   nanoseconds.
2. Using an RTree to filter candidate polygons, instead of iterating through all
   polygons, reduces the execution times of the Ray Casting algorithm.
3. Using a finely-tuned Ray Casting algorithm package
   <https://github.com/tidwall/geojson> to verify whether a polygon contains a
   point.

That's all. There are no black magic tricks inside the tzf package.

The benchmark was conducted using version
<https://github.com/ringsaturn/tzf/releases/tag/v0.10.0>

```
goos: darwin
goarch: amd64
pkg: github.com/ringsaturn/tzf
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkDefaultFinder_GetTimezoneName_Random_WorldCities-16              441309              2778 ns/op              1000 ns/p50            10000 ns/p90            19000 ns/p99
BenchmarkFuzzyFinder_GetTimezoneName_Random_WorldCities-16               1000000              1077 ns/op              1000 ns/p50             2000 ns/p90             2000 ns/p99
BenchmarkGetTimezoneName-16                                               226834              5190 ns/op              5000 ns/p50             5000 ns/p90            22000 ns/p99
BenchmarkGetTimezoneNameAtEdge-16                                         211555              5606 ns/op              5000 ns/p50             6000 ns/p90            23000 ns/p99
BenchmarkGetTimezoneName_Random_WorldCities-16                            163000              7279 ns/op              7000 ns/p50            10000 ns/p90            29000 ns/p99
BenchmarkFullFinder_GetTimezoneName-16                                    212896              5556 ns/op              5000 ns/p50             6000 ns/p90            22000 ns/p99
BenchmarkFullFinder_GetTimezoneNameAtEdge-16                              195381              6262 ns/op              6000 ns/p50             7000 ns/p90            23000 ns/p99
BenchmarkFullFinder_GetTimezoneName_Random_WorldCities-16                 116652              9354 ns/op              8000 ns/p50            15000 ns/p90            31000 ns/p99
PASS
ok      github.com/ringsaturn/tzf       18.321s
```

- <https://ringsaturn.github.io/tzf/> displays continuous benchmarking results.
- <https://ringsaturn.github.io/tz-benchmark/> displays a continuous benchmark
  comparison with other packages.

## Related Repos

- <https://github.com/ringsaturn/tzf-rel> Preprocessed probuf data release repo
- <https://github.com/ringsaturn/tz-benchmark> Continuous Benchmark Compared
  with other packages
- <https://github.com/ringsaturn/tzf-rs> Rust port of tzf
- <https://github.com/ringsaturn/tzfpy> Rust port's Python binding
- <https://github.com/ringsaturn/tzf-server> HTTP&Redis server build with tzf
- <https://github.com/ringsaturn/redizone> Redis compatible server build with
  tzf-rs

## Thanks

- <https://github.com/paulmach/orb>
- <https://github.com/tidwall/geojson>
- <https://github.com/jannikmi/timezonefinder>
- <https://github.com/evansiroky/timezone-boundary-builder>
