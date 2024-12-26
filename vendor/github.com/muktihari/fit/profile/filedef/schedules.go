// Copyright 2024 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// Schedules files are used to schedule a user’s workouts and may contain multiple schedule messages each representing the start time of a workout.
type Schedules struct {
	FileId mesgdef.FileId

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	Schedules []*mesgdef.Schedule

	UnrelatedMessages []proto.Message
}

var _ File = (*Schedules)(nil)

// NewSchedules creates new Schedules File.
func NewSchedules(mesgs ...proto.Message) *Schedules {
	f := &Schedules{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}
	return f
}

// Add adds mesg to the Schedules.
func (f *Schedules) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.Schedule:
		f.Schedules = append(f.Schedules, mesgdef.NewSchedule(&mesg))
	default:
		mesg.Fields = sliceutil.Clone(mesg.Fields)
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

// ToFIT converts Schedules to proto.FIT. If options is nil, default options will be used.
func (f *Schedules) ToFIT(options *mesgdef.Options) proto.FIT {
	var size = 1 // non slice fields

	size += len(f.Schedules) + len(f.DeveloperDataIds) +
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
	for i := range f.Schedules {
		fit.Messages = append(fit.Messages, f.Schedules[i].ToMesg(options))
	}

	// Schedules File does not have fields require sorting,
	// only sort unrelated messages in case it matters.
	SortMessagesByTimestamp(f.UnrelatedMessages)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	return fit
}
