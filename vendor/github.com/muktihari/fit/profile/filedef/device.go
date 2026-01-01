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

// Device files contain information about a device’s file structure/capabilities.
type Device struct {
	FileId mesgdef.FileId // required fields: type, manufacturer, product, serial_number

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	Softwares         []*mesgdef.Software
	Capabilities      []*mesgdef.Capabilities
	FileCapabilities  []*mesgdef.FileCapabilities
	MesgCapabilities  []*mesgdef.MesgCapabilities
	FieldCapabilities []*mesgdef.FieldCapabilities

	UnrelatedMessages []proto.Message
}

var _ File = (*Device)(nil)

// NewDevice creates new Device File.
func NewDevice(mesgs ...proto.Message) *Device {
	f := &Device{FileId: newFileId}
	f.FileId.Type = typedef.FileDevice
	for i := range mesgs {
		f.Add(mesgs[i])
	}
	return f
}

// Add adds mesg to the Device.
func (f *Device) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId.Reset(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.Software:
		f.Softwares = append(f.Softwares, mesgdef.NewSoftware(&mesg))
	case mesgnum.Capabilities:
		f.Capabilities = append(f.Capabilities, mesgdef.NewCapabilities(&mesg))
	case mesgnum.FileCapabilities:
		f.FileCapabilities = append(f.FileCapabilities, mesgdef.NewFileCapabilities(&mesg))
	case mesgnum.MesgCapabilities:
		f.MesgCapabilities = append(f.MesgCapabilities, mesgdef.NewMesgCapabilities(&mesg))
	case mesgnum.FieldCapabilities:
		f.FieldCapabilities = append(f.FieldCapabilities, mesgdef.NewFieldCapabilities(&mesg))
	default:
		mesg.Fields = sliceutil.Clone(mesg.Fields)
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

// ToFIT converts Device to proto.FIT. If options is nil, default options will be used.
func (f *Device) ToFIT(options *mesgdef.Options) proto.FIT {
	var size = 1 // non slice fields

	size += len(f.Softwares) + len(f.Capabilities) + len(f.FileCapabilities) +
		len(f.MesgCapabilities) + len(f.FieldCapabilities) +
		len(f.DeveloperDataIds) + len(f.FieldDescriptions) + len(f.UnrelatedMessages)

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
	for i := range f.Softwares {
		fit.Messages = append(fit.Messages, f.Softwares[i].ToMesg(options))
	}
	for i := range f.Capabilities {
		fit.Messages = append(fit.Messages, f.Capabilities[i].ToMesg(options))
	}
	for i := range f.FileCapabilities {
		fit.Messages = append(fit.Messages, f.FileCapabilities[i].ToMesg(options))
	}
	for i := range f.MesgCapabilities {
		fit.Messages = append(fit.Messages, f.MesgCapabilities[i].ToMesg(options))
	}
	for i := range f.FieldCapabilities {
		fit.Messages = append(fit.Messages, f.FieldCapabilities[i].ToMesg(options))
	}

	// Device File does not have fields require sorting,
	// only sort unrelated messages in case it matters.
	SortMessagesByTimestamp(f.UnrelatedMessages)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	return fit
}
