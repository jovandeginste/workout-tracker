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

// Weight contains time-stamped discrete measurement data of weight.
type Weight struct {
	FileId mesgdef.FileId // required fields: type, manufacturer, product, serial_number

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	UserProfile  *mesgdef.UserProfile
	WeightScales []*mesgdef.WeightScale
	DeviceInfos  []*mesgdef.DeviceInfo

	UnrelatedMessages []proto.Message
}

var _ File = (*Weight)(nil)

// NewWeight creates new Weight File.
func NewWeight(mesgs ...proto.Message) *Weight {
	f := &Weight{FileId: newFileId}
	f.FileId.Type = typedef.FileWeight
	for i := range mesgs {
		f.Add(mesgs[i])
	}
	return f
}

// Add adds mesg to the Weight.
func (f *Weight) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId.Reset(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.UserProfile:
		f.UserProfile = mesgdef.NewUserProfile(&mesg)
	case mesgnum.WeightScale:
		f.WeightScales = append(f.WeightScales, mesgdef.NewWeightScale(&mesg))
	case mesgnum.DeviceInfo:
		f.DeviceInfos = append(f.DeviceInfos, mesgdef.NewDeviceInfo(&mesg))
	default:
		mesg.Fields = sliceutil.Clone(mesg.Fields)
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

// ToFIT converts Weight to proto.FIT. If options is nil, default options will be used.
func (f *Weight) ToFIT(options *mesgdef.Options) proto.FIT {
	var size = 2 // non slice fields

	size += len(f.WeightScales) + len(f.DeviceInfos) +
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
	if f.UserProfile != nil {
		fit.Messages = append(fit.Messages, f.UserProfile.ToMesg(options))
	}
	for i := range f.WeightScales {
		fit.Messages = append(fit.Messages, f.WeightScales[i].ToMesg(options))
	}
	for i := range f.DeviceInfos {
		fit.Messages = append(fit.Messages, f.DeviceInfos[i].ToMesg(options))
	}

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages[sortStartPos:])

	return fit
}
