// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sliceutil contains helper functions that either slices package does not
// provide or it's provided but not suitable for our use case.
package sliceutil

// Clone shallow copies s, but unlike slices.Clone that uses append variant,
// this uses make+copy variant so we don't generate additional unused capacity.
// This will return nil when len(s) == 0 instead of returning zerobase slice
// since we use it to clone a slice that is backed by an array pool.
//
// Related issues regarding slices.Clone that we are trying to solve with this:
//   - https://go.dev/issue/68488: keep alive underlying array on zero len slice,
//     affected Go version: v1.22.0 up to v.1.23.X. CL has been merged but it seems
//     without a backport, so the issue might not be resolved until v1.24.0.
//   - https://go.dev/issue/53643: suboptimal performance
func Clone[S ~[]E, E any](s S) (s2 S) {
	if len(s) == 0 {
		return nil
	}
	s2 = make(S, len(s))
	copy(s2, s)
	return s2
}
