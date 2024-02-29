package fit

import (
	"bytes"
	"strconv"
	"sync"
)

type jsonEncodeState struct {
	bytes.Buffer
	scratch [64]byte
}

var jsonEncodeStatePool sync.Pool

func newJSONEncodeState() *jsonEncodeState {
	if v := jsonEncodeStatePool.Get(); v != nil {
		e := v.(*jsonEncodeState)
		e.Reset()
		return e
	}
	return new(jsonEncodeState)
}

func (j *jsonEncodeState) writeFieldName(s string) {
	j.WriteByte('"')
	j.WriteString(s)
	j.WriteByte('"')
	j.WriteByte(':')
}

func (j *jsonEncodeState) writeStringBytes(b []byte) {
	j.WriteByte('"')
	j.Write(b)
	j.WriteByte('"')
}

func (j *jsonEncodeState) writeUint(u uint64) {
	b := strconv.AppendUint(j.scratch[:0], u, 10)
	j.Write(b)
}

func (j *jsonEncodeState) open() {
	j.WriteByte('{')
}

func (j *jsonEncodeState) close() {
	j.WriteByte('}')
}

func (j *jsonEncodeState) c() {
	j.WriteByte(',')
}
