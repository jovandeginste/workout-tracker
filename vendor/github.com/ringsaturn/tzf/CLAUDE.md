# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

TZF is a high-performance timezone finder library for Go that determines the timezone for given latitude and longitude coordinates. The project is designed for geospatial services like weather forecast APIs where fast timezone lookups are critical.

## Core Architecture

The project implements three main finder types with different performance/accuracy trade-offs:

### Finder Types
- **Finder** (`tzf.go`): Polygon-based finder using point-in-polygon algorithms with RTree spatial indexing. Most accurate but memory intensive (~100MB lite, ~1GB full data).
- **FuzzyFinder** (`tzf_fuzzy.go`): Tile-based finder using pre-indexed map tiles for ultra-fast lookups (~1.78MB data). Slightly less accurate but extremely fast.
- **DefaultFinder** (`tzf_default_finder.go`): Hybrid approach combining FuzzyFinder and compressed Finder. Uses FuzzyFinder first, falls back to Finder with spatial tolerance for edge cases.

### Data Pipeline
The project processes timezone data through several stages:
1. Raw GeoJSON from timezone-boundary-builder → Full protobuf data (cmd/geojson2tzpb)
2. Full data → Reduced/Lite data (cmd/reducetzpb) 
3. Lite data → Compressed data (cmd/compresstzpb)
4. Lite data → Pre-indexed tiles (cmd/preindextzpb)

### Key Components
- **Interface F** (`f.go`): Common interface for all finder implementations
- **Convert package** (`convert/`): Handles GeoJSON to protobuf conversion
- **Reduce package** (`reduce/`): Data compression and decompression
- **Preindex package** (`preindex/`): Tile-based pre-indexing
- **Protocol Buffers** (`pb/tzf/v1/`): Data serialization format

## Development Commands

### Build and Testing
```bash
# Format code
make fmt

# Run tests with linting and coverage
make test

# Generate coverage report
make cover

# Run benchmarks  
make bench

# Generate protocol buffers
make pb
```

### Key Build Tools
- **golangci-lint**: Linting (required for tests)
- **buf**: Protocol buffer generation
- **go-licenses**: Third-party license management

### CLI Tool
The project includes a CLI tool at `cmd/tzf/main.go`:
```bash
go install github.com/ringsaturn/tzf/cmd/tzf@latest
tzf -lng 116.3883 -lat 39.9289
```

## Data Management

The project uses external data from `github.com/ringsaturn/tzf-rel-lite` embedded as Go modules. Data versions must match between Finder and FuzzyFinder components.

### Memory Usage Patterns
- DefaultFinder: ~150MB init, ~60MB after GC
- Full accuracy Finder: ~900MB init, ~660MB after GC
- FuzzyFinder: ~1.78MB minimal memory usage

## Performance Characteristics

The codebase is optimized for sub-microsecond timezone lookups:
- Pre-indexing handles most queries in ~1000ns
- RTree spatial indexing reduces polygon iteration
- Optimized ray-casting algorithm via tidwall/geojson
- Fallback mechanisms for edge cases with spatial tolerance

## Testing and Validation

Tests should verify:
- Timezone lookup accuracy for world cities
- Performance benchmarks across finder types
- Data version compatibility
- Memory usage patterns
- Edge case handling (borders, oceans)