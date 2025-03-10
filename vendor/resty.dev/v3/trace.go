// Copyright (c) 2015-present Jeevanandam M (jeeva@myjeeva.com), All rights reserved.
// resty source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package resty

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http/httptrace"
	"time"
)

// TraceInfo struct is used to provide request trace info such as DNS lookup
// duration, Connection obtain duration, Server processing duration, etc.
type TraceInfo struct {
	// DNSLookup is the duration that transport took to perform
	// DNS lookup.
	DNSLookup time.Duration `json:"dns_lookup_time"`

	// ConnTime is the duration it took to obtain a successful connection.
	ConnTime time.Duration `json:"connection_time"`

	// TCPConnTime is the duration it took to obtain the TCP connection.
	TCPConnTime time.Duration `json:"tcp_connection_time"`

	// TLSHandshake is the duration of the TLS handshake.
	TLSHandshake time.Duration `json:"tls_handshake_time"`

	// ServerTime is the server's duration for responding to the first byte.
	ServerTime time.Duration `json:"server_time"`

	// ResponseTime is the duration since the first response byte from the server to
	// request completion.
	ResponseTime time.Duration `json:"response_time"`

	// TotalTime is the duration of the total time request taken end-to-end.
	TotalTime time.Duration `json:"total_time"`

	// IsConnReused is whether this connection has been previously
	// used for another HTTP request.
	IsConnReused bool `json:"is_connection_reused"`

	// IsConnWasIdle is whether this connection was obtained from an
	// idle pool.
	IsConnWasIdle bool `json:"is_connection_was_idle"`

	// ConnIdleTime is the duration how long the connection that was previously
	// idle, if IsConnWasIdle is true.
	ConnIdleTime time.Duration `json:"connection_idle_time"`

	// RequestAttempt is to represent the request attempt made during a Resty
	// request execution flow, including retry count.
	RequestAttempt int `json:"request_attempt"`

	// RemoteAddr returns the remote network address.
	RemoteAddr string `json:"remote_address"`
}

// String method returns string representation of request trace information.
func (ti TraceInfo) String() string {
	return fmt.Sprintf(`TRACE INFO:
  DNSLookupTime : %v
  ConnTime      : %v
  TCPConnTime   : %v
  TLSHandshake  : %v
  ServerTime    : %v
  ResponseTime  : %v
  TotalTime     : %v
  IsConnReused  : %v
  IsConnWasIdle : %v
  ConnIdleTime  : %v
  RequestAttempt: %v
  RemoteAddr    : %v`, ti.DNSLookup, ti.ConnTime, ti.TCPConnTime,
		ti.TLSHandshake, ti.ServerTime, ti.ResponseTime, ti.TotalTime,
		ti.IsConnReused, ti.IsConnWasIdle, ti.ConnIdleTime, ti.RequestAttempt,
		ti.RemoteAddr)
}

// JSON method returns the JSON string of request trace information
func (ti TraceInfo) JSON() string {
	return toJSON(ti)
}

// Clone method returns the clone copy of [TraceInfo]
func (ti TraceInfo) Clone() *TraceInfo {
	ti2 := new(TraceInfo)
	*ti2 = ti
	return ti2
}

// clientTrace struct maps the [httptrace.ClientTrace] hooks into Fields
// with the same naming for easy understanding. Plus additional insights
// [Request].
type clientTrace struct {
	getConn              time.Time
	dnsStart             time.Time
	dnsDone              time.Time
	connectDone          time.Time
	tlsHandshakeStart    time.Time
	tlsHandshakeDone     time.Time
	gotConn              time.Time
	gotFirstResponseByte time.Time
	endTime              time.Time
	gotConnInfo          httptrace.GotConnInfo
}

func (t *clientTrace) createContext(ctx context.Context) context.Context {
	return httptrace.WithClientTrace(
		ctx,
		&httptrace.ClientTrace{
			DNSStart: func(_ httptrace.DNSStartInfo) {
				t.dnsStart = time.Now()
			},
			DNSDone: func(_ httptrace.DNSDoneInfo) {
				t.dnsDone = time.Now()
			},
			ConnectStart: func(_, _ string) {
				if t.dnsDone.IsZero() {
					t.dnsDone = time.Now()
				}
				if t.dnsStart.IsZero() {
					t.dnsStart = t.dnsDone
				}
			},
			ConnectDone: func(net, addr string, err error) {
				t.connectDone = time.Now()
			},
			GetConn: func(_ string) {
				t.getConn = time.Now()
			},
			GotConn: func(ci httptrace.GotConnInfo) {
				t.gotConn = time.Now()
				t.gotConnInfo = ci
			},
			GotFirstResponseByte: func() {
				t.gotFirstResponseByte = time.Now()
			},
			TLSHandshakeStart: func() {
				t.tlsHandshakeStart = time.Now()
			},
			TLSHandshakeDone: func(_ tls.ConnectionState, _ error) {
				t.tlsHandshakeDone = time.Now()
			},
		},
	)
}
