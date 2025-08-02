// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/profile/factory"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
)

// Factory defines a contract that any Factory containing these method can be used by mesgdef's structs.
type Factory interface {
	// CreateField creates new field based on defined messages in the factory.
	// If not found, it returns new field with "unknown" name.
	CreateField(mesgNum typedef.MesgNum, num byte) proto.Field
}

type Options struct {
	Factory               Factory // If not specified, factory.StandardFactory() will be used.
	IncludeExpandedFields bool
}

var defaultOptions = DefaultOptions()

func DefaultOptions() *Options {
	return &Options{
		Factory:               factory.StandardFactory(),
		IncludeExpandedFields: false,
	}
}
