// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// Totals files are used to summarize a user’s activities and may contain multiple totals messages each representing
// summaries of a different activity type/sport
type Totals struct {
	FileId mesgdef.FileId

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	Totals []*mesgdef.Totals

	UnrelatedMessages []proto.Message
}

var _ File = (*Totals)(nil)

// NewTotals creates new Totals File.
func NewTotals(mesgs ...proto.Message) *Totals {
	f := &Totals{FileId: newFileId}
	f.FileId.Type = typedef.FileTotals
	for i := range mesgs {
		f.Add(mesgs[i])
	}
	return f
}

// Add adds mesg to the Totals.
func (f *Totals) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId.Reset(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.Totals:
		f.Totals = append(f.Totals, mesgdef.NewTotals(&mesg))
	default:
		mesg.Fields = sliceutil.Clone(mesg.Fields)
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

// ToFIT converts Totals to proto.FIT. If options is nil, default options will be used.
func (f *Totals) ToFIT(options *mesgdef.Options) proto.FIT {
	var size = 1 // non slice fields

	size += len(f.Totals) + len(f.DeveloperDataIds) +
		len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.FIT{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	var sortStartPos = 1 + len(f.DeveloperDataIds) + len(f.FieldDescriptions)
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(options))

	for i := range f.DeveloperDataIds {
		fit.Messages = append(fit.Messages, f.DeveloperDataIds[i].ToMesg(options))
	}
	for i := range f.FieldDescriptions {
		fit.Messages = append(fit.Messages, f.FieldDescriptions[i].ToMesg(options))
	}
	for i := range f.Totals {
		fit.Messages = append(fit.Messages, f.Totals[i].ToMesg(options))
	}

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages[sortStartPos:])

	return fit
}
