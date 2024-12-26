// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package filedef contains general implementation of known common file types retrieved from
// Garmin or its affiliates website and a listener building block convert decoded FIT file into
// the desired common file type as soon as the message is decoded.
//
// You may find that the common file types declared here are not sufficient for your need,
// but don't worry, you can always create your own common file types that suit your need more;
// whether creating a fresh one or embedding the existing; and still be able to use listener
// building block as long as it satisfy the File interface.
package filedef
