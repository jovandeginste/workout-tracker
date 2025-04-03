# [tzf](https://github.com/ringsaturn/tzf)'s preprocessed timezone data

![](https://img.shields.io/github/v/release/ringsaturn/tzf-rel-lite?style=plastic)

## Update Data Steps

### CI

Data build in [GitHub Actions](.github/workflows/ci.yml).

### Local build

1. Install CLI tool

```bash
# install tools
go install github.com/ringsaturn/tzf/cmd/geojson2tzpb@latest
go install github.com/ringsaturn/tzf/cmd/reducetzpb@latest
go install github.com/ringsaturn/tzf/cmd/compresstzpb@latest
go install github.com/ringsaturn/tzf/cmd/preindextzpb@latest
```

2. Set data version to build(Below steps need this environment var)

```bash
export TIMEZONE_BOUNDARY_VERSION=2023b
```

3. Download data

```bash
# download data
wget https://github.com/evansiroky/timezone-boundary-builder/releases/download/${TIMEZONE_BOUNDARY_VERSION}/timezones-with-oceans.geojson.zip
unzip timezones-with-oceans.geojson.zip
```

4. Make data

```bash
geojson2tzpb combined-with-oceans.json | xargs reducetzpb | xargs compresstzpb
preindextzpb combined-with-oceans.reduce.bin
```

## References

- Protocol Buffers define:
  <https://github.com/ringsaturn/tzf/blob/main/pb/tzinfo.proto>
- Maintain tools
  - [`checkboundaryrelease`](https://github.com/ringsaturn/tzf/tree/main/cmd/checkboundaryrelease)
  - [`geojson2tzpb`](https://github.com/ringsaturn/tzf/tree/main/cmd/geojson2tzpb)
  - [`reducetzpb`](https://github.com/ringsaturn/tzf/tree/main/cmd/reducetzpb)
  - [`compresstzpb`](https://github.com/ringsaturn/tzf/tree/main/cmd/compresstzpb)
  - [`preindextzpb`](https://github.com/ringsaturn/tzf/tree/main/cmd/preindextzpb)
- To view data in GeoJSON format(which is more convenient to show on map), see
  <https://github.com/ringsaturn/tzf-server>

- Data Usage examples
  - Go: <https://github.com/ringsaturn/tzf>
  - Rust: <https://github.com/ringsaturn/tzf-rs>
