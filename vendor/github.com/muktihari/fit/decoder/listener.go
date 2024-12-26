// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decoder

import "github.com/muktihari/fit/proto"

// MesgListener is an interface for listening to message decoded events.
type MesgListener interface {
	// OnMesg receives message from Decoder. The lifecycle of the mesg object is only guaranteed before
	// OnMesg returns. Any listener that wants to process the mesg concurrently should copy the mesg,
	// otherwise, the value of the mesg might be changed by the time it is being processed. Except,
	// the Decoder is directed to copy the mesg before passing the mesg to listener using Option.
	OnMesg(mesg proto.Message)
}

// MesgDefListener is an interface for listening to message definition decoded event.
type MesgDefListener interface {
	// OnMesgDef receives message definition from Decoder.
	OnMesgDef(mesgDef proto.MessageDefinition)
}
