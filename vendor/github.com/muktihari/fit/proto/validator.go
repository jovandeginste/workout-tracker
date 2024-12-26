// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"fmt"

	"github.com/muktihari/fit/profile/basetype"
)

const ErrProtocolViolation = errorString("protocol violation")

// NewValidator creates protocol validator base on given version.
func NewValidator(protocolVersion Version) *Validator {
	return &Validator{ProtocolVersion: protocolVersion}
}

// Validator is protocol validator
type Validator struct{ ProtocolVersion Version }

// ValidateMessageDefinition validates whether the message definition contains unsupported data for the targeted version.
func (p *Validator) ValidateMessageDefinition(mesgDef *MessageDefinition) error {
	if p.ProtocolVersion == V1 {
		if len(mesgDef.DeveloperFieldDefinitions) > 0 {
			return fmt.Errorf("protocol version 1.0 do not support developer fields: %w", ErrProtocolViolation)
		}
		for _, fieldDef := range mesgDef.FieldDefinitions {
			if fieldDef.BaseType&basetype.BaseTypeNumMask > basetype.Byte&basetype.BaseTypeNumMask { // byte was the last type added in 1.0
				return fmt.Errorf("protocol version 1.0 do not support type '%s': %w", fieldDef.BaseType, ErrProtocolViolation)
			}
		}
		return nil
	}
	return nil
}

// ValidateMessage validates whether the message contains unsupported data for the targeted version.
func (p *Validator) ValidateMessage(mesg *Message) error {
	if p.ProtocolVersion == V1 {
		if len(mesg.DeveloperFields) > 0 {
			return fmt.Errorf("protocol version 1.0 do not support developer fields: %w", ErrProtocolViolation)
		}
		for i := range mesg.Fields {
			field := &mesg.Fields[i]
			if field.BaseType&basetype.BaseTypeNumMask > basetype.Byte&basetype.BaseTypeNumMask { // byte was the last type added in 1.0
				return fmt.Errorf("protocol version 1.0 do not support type '%s': %w", field.BaseType, ErrProtocolViolation)
			}
		}
	}
	return nil
}
