// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"fmt"

	"github.com/muktihari/fit/profile/basetype"
)

const ErrProtocolViolation = errorString("protocol violation")

// Validator is protocol validator
var Validator validator

type validator struct{}

// ValidateMessageDefinition validates whether the message definition contains unsupported data for the targeted version.
func (validator) ValidateMessageDefinition(mesgDef *MessageDefinition, v Version) error {
	if v == V1 {
		if len(mesgDef.DeveloperFieldDefinitions) > 0 {
			return fmt.Errorf("protocol version 1.0 do not support developer fields: %w", ErrProtocolViolation)
		}
		for _, fieldDef := range mesgDef.FieldDefinitions {
			if fieldDef.BaseType&basetype.BaseTypeNumMask > basetype.Byte&basetype.BaseTypeNumMask { // byte was the last type added in 1.0
				return fmt.Errorf("protocol version 1.0 do not support type %q: %w", fieldDef.BaseType, ErrProtocolViolation)
			}
		}
		return nil
	}
	return nil
}

// ValidateMessage validates whether the message contains unsupported data for the targeted version.
func (validator) ValidateMessage(mesg *Message, v Version) error {
	if v == V1 {
		if len(mesg.DeveloperFields) > 0 {
			return fmt.Errorf("protocol version 1.0 do not support developer fields: %w", ErrProtocolViolation)
		}
		for i := range mesg.Fields {
			field := &mesg.Fields[i]
			if field.BaseType&basetype.BaseTypeNumMask > basetype.Byte&basetype.BaseTypeNumMask { // byte was the last type added in 1.0
				return fmt.Errorf("protocol version 1.0 do not support type %q: %w", field.BaseType, ErrProtocolViolation)
			}
		}
	}
	return nil
}
