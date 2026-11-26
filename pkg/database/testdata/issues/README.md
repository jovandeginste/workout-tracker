# Donated workout samples

These files were donated through the GitHub issue tracker of this repository.
They are fixtures for the import pipeline (`converters.Parse` and
`database.NewWorkout`). Each sample has a golden file with the same name plus
`.json`, holding the expected parse result. Per-sample provenance is recorded
in `sources.json`.

Some files were attached with a non-workout extension; they are stored with the
extension that matches their content. Archives were unpacked, and identical
files were kept once. Partial or redacted snippets pasted inline are not
included.

Regenerate the golden files after an intentional parser change:

```sh
go test ./pkg/database -run TestIssueSamples -count=1 -update-issue-samples
```

The golden files keep the workout summary, detected climbs and descents, the
number of parsed points, a SHA-256 digest of every point (including slope
grades), the first and last point, and a digest of the derived heatmap points.
Derived values that the background worker normally calculates (slope grades,
climb and descent detection, cadence and power averages) are included. Storing
all points would make the goldens larger than the samples. A digest mismatch
means the point parsing changed even when the summary looks the same.
