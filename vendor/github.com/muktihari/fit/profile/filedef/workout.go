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

// Workout is a file contains instructions for performing a structured activity.
//
// ref: https://developer.garmin.com/fit/file-types/workout/
type Workout struct {
	FileId mesgdef.FileId // required fields: type, manufacturer, product, serial_number, time_created

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	// Required Messages
	Workout      *mesgdef.Workout       // required fields: num_valid_steps
	WorkoutSteps []*mesgdef.WorkoutStep // required fields: message_index, duration_type, target_type

	// Messages not related to Workout
	UnrelatedMessages []proto.Message
}

var _ File = (*Workout)(nil)

// NewWorkout creates new Workout File.
func NewWorkout(mesgs ...proto.Message) *Workout {
	f := &Workout{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}
	return f
}

// Add adds mesg to the Workout.
func (f *Workout) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.Workout:
		f.Workout = mesgdef.NewWorkout(&mesg)
	case mesgnum.WorkoutStep:
		f.WorkoutSteps = append(f.WorkoutSteps, mesgdef.NewWorkoutStep(&mesg))
	default:
		mesg.Fields = sliceutil.Clone(mesg.Fields)
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

// ToFIT converts Workout to proto.FIT. If options is nil, default options will be used.
func (f *Workout) ToFIT(options *mesgdef.Options) proto.FIT {
	size := 2 /* non slice fields */

	size += len(f.WorkoutSteps) + len(f.DeveloperDataIds) + len(f.FieldDescriptions) + len(f.UnrelatedMessages)

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
	if f.Workout != nil {
		fit.Messages = append(fit.Messages, f.Workout.ToMesg(options))
	}
	for i := range f.WorkoutSteps {
		fit.Messages = append(fit.Messages, f.WorkoutSteps[i].ToMesg(options))
	}

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	return fit
}
