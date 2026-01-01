// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedef

import (
	"slices"

	"github.com/muktihari/fit/profile/mesgdef"
	"github.com/muktihari/fit/profile/untyped/fieldnum"
	"github.com/muktihari/fit/profile/untyped/mesgnum"
	"github.com/muktihari/fit/proto"
)

var newFileId = *mesgdef.NewFileId(nil)

// File is an interface for defining common type file, any defined common file type should implement
// the following methods to be able to work with Listener (and other building block in filedef package).
type File interface {
	// Add adds message into file structure.
	Add(mesg proto.Message)
	// ToFIT converts file back to proto.FIT structure.
	ToFIT(options *mesgdef.Options) proto.FIT
}

// SortMessagesByTimestamp sorts messages by timestamp. The following rules will apply:
//   - Any message without timestamp field will be placed to the beginning of the slice
//     to enable these messages to be retrieved early such as UserProfile.
//   - Any message with invalid timestamp will be places at the end of the slices.
//
// Special Case:
//
// All timestamp fields should have num 253, except:
//   - Course Point's Timestamp num: 1
//   - Set's Timestamp num: 254
//
// We will sort these timestamps accordingly since the messages' order matters.
//
// For details, see [github.com/muktihari/fit/proto.FieldNumTimestamp] doc.
func SortMessagesByTimestamp(messages []proto.Message) {
	slices.SortStableFunc(messages, func(m1, m2 proto.Message) int {
		var f1, f2 *proto.Field
		switch m1.Num {
		case mesgnum.CoursePoint:
			f1 = m1.FieldByNum(fieldnum.CoursePointTimestamp)
		case mesgnum.Set:
			f1 = m1.FieldByNum(fieldnum.SetTimestamp)
		default:
			f1 = m1.FieldByNum(proto.FieldNumTimestamp)
		}

		switch m2.Num {
		case mesgnum.CoursePoint:
			f2 = m2.FieldByNum(fieldnum.CoursePointTimestamp)
		case mesgnum.Set:
			f2 = m2.FieldByNum(fieldnum.SetTimestamp)
		default:
			f2 = m2.FieldByNum(proto.FieldNumTimestamp)
		}

		// Place messages which does not have a timestamp at the beginning of the slice.
		if f1 == nil && f2 == nil {
			return 0
		} else if f1 == nil {
			return -1
		} else if f2 == nil {
			return 1
		}

		// Sort timestamps regardless of whether any of the values are invalid.
		// Any invalid value will be placed at the end of the slice.
		t1 := f1.Value.Uint32()
		t2 := f2.Value.Uint32()
		if t1 < t2 {
			return -1
		}
		if t1 > t2 {
			return 1
		}
		return 0
	})
}
