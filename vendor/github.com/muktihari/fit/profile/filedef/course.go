// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

// Course is a common file type used as points of courses to assist with on- and off-road navigation,
// to provide turn by turn directions, or with virtual training applications to simulate real-world activities.
//
// Please note since we group the same mesgdef types in slices, we lose the arrival order of the messages.
// But for messages that have timestamp, we can reconstruct the messages by timestamp order.
//
// ref: https://developer.garmin.com/fit/file-types/course/
type Course struct {
	FileId mesgdef.FileId // must have mesg

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	// Required Messages
	Course  *mesgdef.Course
	Lap     *mesgdef.Lap
	Records []*mesgdef.Record
	Events  []*mesgdef.Event

	// Optional Messages
	CoursePoints []*mesgdef.CoursePoint

	// Messages not related to Course
	UnrelatedMessages []proto.Message
}

var _ File = (*Course)(nil)

// NewCourse creates new Course File.
func NewCourse(mesgs ...proto.Message) *Course {
	f := &Course{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}
	return f
}

// Add adds mesg to the Course.
func (f *Course) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.Course:
		f.Course = mesgdef.NewCourse(&mesg)
	case mesgnum.Lap:
		f.Lap = mesgdef.NewLap(&mesg)
	case mesgnum.Record:
		f.Records = append(f.Records, mesgdef.NewRecord(&mesg))
	case mesgnum.Event:
		f.Events = append(f.Events, mesgdef.NewEvent(&mesg))
	case mesgnum.CoursePoint:
		f.CoursePoints = append(f.CoursePoints, mesgdef.NewCoursePoint(&mesg))
	default:
		mesg.Fields = sliceutil.Clone(mesg.Fields)
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

// ToFIT converts Course to proto.FIT. If options is nil, default options will be used.
func (f *Course) ToFIT(options *mesgdef.Options) proto.FIT {
	size := 3 /* non slice fields */

	size += len(f.Records) + len(f.Events) + len(f.CoursePoints) +
		len(f.DeveloperDataIds) + len(f.FieldDescriptions) + len(f.UnrelatedMessages)

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
	if f.Course != nil {
		fit.Messages = append(fit.Messages, f.Course.ToMesg(options))
	}
	if f.Lap != nil {
		fit.Messages = append(fit.Messages, f.Lap.ToMesg(options))
	}
	for i := range f.Records {
		fit.Messages = append(fit.Messages, f.Records[i].ToMesg(options))
	}
	for i := range f.Events {
		fit.Messages = append(fit.Messages, f.Events[i].ToMesg(options))
	}
	for i := range f.CoursePoints {
		fit.Messages = append(fit.Messages, f.CoursePoints[i].ToMesg(options))
	}

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages[sortStartPos:])

	return fit
}
