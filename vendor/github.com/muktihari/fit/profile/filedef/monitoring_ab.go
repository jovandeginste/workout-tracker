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

// MonitoringAB (Monitoring A and Monitoring B) files are used to store data that is logged over varying time intervals.
// The two monitoring file formats are identical apart from supporting different conventions for file_id.number and the start of accumulating data values.
//
// The FIT file_id.type = 15 for a monitoring_a file and
// the FIT file_id.type = 32 for a monitoring_b file
type MonitoringAB struct {
	FileId mesgdef.FileId

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	MonitoringInfo *mesgdef.MonitoringInfo
	Monitorings    []*mesgdef.Monitoring
	DeviceInfos    []*mesgdef.DeviceInfo

	UnrelatedMessages []proto.Message
}

var _ File = (*MonitoringAB)(nil)

// NewMonitoringAB creates new MonitoringAB File.
func NewMonitoringAB(mesgs ...proto.Message) *MonitoringAB {
	f := &MonitoringAB{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}
	return f
}

// Add adds mesg to the MonitoringAB.
func (f *MonitoringAB) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.MonitoringInfo:
		f.MonitoringInfo = mesgdef.NewMonitoringInfo(&mesg)
	case mesgnum.Monitoring:
		f.Monitorings = append(f.Monitorings, mesgdef.NewMonitoring(&mesg))
	case mesgnum.DeviceInfo:
		f.DeviceInfos = append(f.DeviceInfos, mesgdef.NewDeviceInfo(&mesg))
	default:
		mesg.Fields = sliceutil.Clone(mesg.Fields)
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

// ToFIT converts MonitoringAB to proto.FIT. If options is nil, default options will be used.
func (f *MonitoringAB) ToFIT(options *mesgdef.Options) proto.FIT {
	var size = 2 // non slice fields

	size += len(f.Monitorings) + len(f.DeviceInfos) +
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
	if f.MonitoringInfo != nil {
		fit.Messages = append(fit.Messages, f.MonitoringInfo.ToMesg(options))
	}
	for i := range f.Monitorings {
		fit.Messages = append(fit.Messages, f.Monitorings[i].ToMesg(options))
	}
	for i := range f.DeviceInfos {
		fit.Messages = append(fit.Messages, f.DeviceInfos[i].ToMesg(options))
	}

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	SortMessagesByTimestamp(fit.Messages[sortStartPos:])

	return fit
}
