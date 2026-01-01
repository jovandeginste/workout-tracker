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

// Segment files contain data defining a route and timing information to gauge progress against previous performances or other users
type Segment struct {
	FileId mesgdef.FileId

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	SegmentId                 *mesgdef.SegmentId
	SegmentLap                *mesgdef.SegmentLap
	SegmentLeaderboardEntries []*mesgdef.SegmentLeaderboardEntry
	SegmentPoints             []*mesgdef.SegmentPoint

	UnrelatedMessages []proto.Message
}

var _ File = (*Segment)(nil)

// NewSegment creates new Segment File.
func NewSegment(mesgs ...proto.Message) *Segment {
	f := &Segment{FileId: newFileId}
	f.FileId.Type = typedef.FileSegment
	for i := range mesgs {
		f.Add(mesgs[i])
	}
	return f
}

// Add adds mesg to the Segment.
func (f *Segment) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId.Reset(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.SegmentId:
		f.SegmentId = mesgdef.NewSegmentId(&mesg)
	case mesgnum.SegmentLeaderboardEntry:
		f.SegmentLeaderboardEntries = append(f.SegmentLeaderboardEntries, mesgdef.NewSegmentLeaderboardEntry(&mesg))
	case mesgnum.SegmentLap:
		f.SegmentLap = mesgdef.NewSegmentLap(&mesg)
	case mesgnum.SegmentPoint:
		f.SegmentPoints = append(f.SegmentPoints, mesgdef.NewSegmentPoint(&mesg))
	default:
		mesg.Fields = sliceutil.Clone(mesg.Fields)
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

// ToFIT converts Segment to proto.FIT. If options is nil, default options will be used.
func (f *Segment) ToFIT(options *mesgdef.Options) proto.FIT {
	var size = 3 // non slice fields

	size += len(f.SegmentPoints) + len(f.SegmentLeaderboardEntries) + len(f.DeveloperDataIds) +
		len(f.FieldDescriptions) + len(f.UnrelatedMessages)

	fit := proto.FIT{
		Messages: make([]proto.Message, 0, size),
	}

	// Should be as ordered: FieldId, DeveloperDataId and FieldDescription
	fit.Messages = append(fit.Messages, f.FileId.ToMesg(options))

	for i := range f.DeveloperDataIds {
		fit.Messages = append(fit.Messages, f.DeveloperDataIds[i].ToMesg(options))
	}
	for i := range f.FieldDescriptions {
		fit.Messages = append(fit.Messages, f.FieldDescriptions[i].ToMesg(options))
	}
	if f.SegmentId != nil {
		fit.Messages = append(fit.Messages, f.SegmentId.ToMesg(options))
	}
	for i := range f.SegmentLeaderboardEntries {
		fit.Messages = append(fit.Messages, f.SegmentLeaderboardEntries[i].ToMesg(options))
	}
	if f.SegmentLap != nil {
		fit.Messages = append(fit.Messages, f.SegmentLap.ToMesg(options))
	}
	for i := range f.SegmentPoints {
		fit.Messages = append(fit.Messages, f.SegmentPoints[i].ToMesg(options))
	}

	// Segment File does not have fields require sorting,
	// only sort unrelated messages in case it matters.
	SortMessagesByTimestamp(f.UnrelatedMessages)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	return fit
}
