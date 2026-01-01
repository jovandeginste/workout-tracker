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

// Sport files contain information about the user’s desired target zones.
type Sport struct {
	FileId mesgdef.FileId // required fields: type, manufacturer, product, serial_number

	// Developer Data Lookup
	DeveloperDataIds  []*mesgdef.DeveloperDataId
	FieldDescriptions []*mesgdef.FieldDescription

	ZonesTargets []*mesgdef.ZonesTarget
	Sport        *mesgdef.Sport
	HrZones      []*mesgdef.HrZone
	PowerZones   []*mesgdef.PowerZone
	MetZones     []*mesgdef.MetZone
	SpeedZones   []*mesgdef.SpeedZone
	CadenceZones []*mesgdef.CadenceZone

	UnrelatedMessages []proto.Message
}

var _ File = (*Sport)(nil)

// NewSport creates new Sport File.
func NewSport(mesgs ...proto.Message) *Sport {
	f := &Sport{}
	for i := range mesgs {
		f.Add(mesgs[i])
	}
	return f
}

// Add adds mesg to the Sport.
func (f *Sport) Add(mesg proto.Message) {
	switch mesg.Num {
	case mesgnum.FileId:
		f.FileId = *mesgdef.NewFileId(&mesg)
	case mesgnum.DeveloperDataId:
		f.DeveloperDataIds = append(f.DeveloperDataIds, mesgdef.NewDeveloperDataId(&mesg))
	case mesgnum.FieldDescription:
		f.FieldDescriptions = append(f.FieldDescriptions, mesgdef.NewFieldDescription(&mesg))
	case mesgnum.ZonesTarget:
		f.ZonesTargets = append(f.ZonesTargets, mesgdef.NewZonesTarget(&mesg))
	case mesgnum.Sport:
		f.Sport = mesgdef.NewSport(&mesg)
	case mesgnum.HrZone:
		f.HrZones = append(f.HrZones, mesgdef.NewHrZone(&mesg))
	case mesgnum.PowerZone:
		f.PowerZones = append(f.PowerZones, mesgdef.NewPowerZone(&mesg))
	case mesgnum.MetZone:
		f.MetZones = append(f.MetZones, mesgdef.NewMetZone(&mesg))
	case mesgnum.SpeedZone:
		f.SpeedZones = append(f.SpeedZones, mesgdef.NewSpeedZone(&mesg))
	case mesgnum.CadenceZone:
		f.CadenceZones = append(f.CadenceZones, mesgdef.NewCadenceZone(&mesg))
	default:
		mesg.Fields = sliceutil.Clone(mesg.Fields)
		f.UnrelatedMessages = append(f.UnrelatedMessages, mesg)
	}
}

// ToFIT converts Sport to proto.FIT. If options is nil, default options will be used.
func (f *Sport) ToFIT(options *mesgdef.Options) proto.FIT {
	var size = 2 // non slice fields

	size += len(f.ZonesTargets) + len(f.HrZones) + len(f.PowerZones) +
		len(f.MetZones) + len(f.SpeedZones) + len(f.CadenceZones) +
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
	for i := range f.ZonesTargets {
		fit.Messages = append(fit.Messages, f.ZonesTargets[i].ToMesg(options))
	}
	if f.Sport != nil {
		fit.Messages = append(fit.Messages, f.Sport.ToMesg(options))
	}
	for i := range f.HrZones {
		fit.Messages = append(fit.Messages, f.HrZones[i].ToMesg(options))
	}
	for i := range f.PowerZones {
		fit.Messages = append(fit.Messages, f.PowerZones[i].ToMesg(options))
	}
	for i := range f.MetZones {
		fit.Messages = append(fit.Messages, f.MetZones[i].ToMesg(options))
	}
	for i := range f.SpeedZones {
		fit.Messages = append(fit.Messages, f.SpeedZones[i].ToMesg(options))
	}
	for i := range f.CadenceZones {
		fit.Messages = append(fit.Messages, f.CadenceZones[i].ToMesg(options))
	}

	// Sport File does not have fields require sorting,
	// only sort unrelated messages in case it matters.
	SortMessagesByTimestamp(f.UnrelatedMessages)

	fit.Messages = append(fit.Messages, f.UnrelatedMessages...)

	return fit
}
