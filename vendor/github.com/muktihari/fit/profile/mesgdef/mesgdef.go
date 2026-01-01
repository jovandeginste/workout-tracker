// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

type Options struct {
	IncludeExpandedFields bool
}

var defaultOptions = DefaultOptions()

func DefaultOptions() *Options {
	return &Options{
		IncludeExpandedFields: false,
	}
}
