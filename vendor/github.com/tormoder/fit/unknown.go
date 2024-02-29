package fit

type unknownField struct {
	mesgNum  MesgNum
	fieldNum byte
}

// UnknownField represents an unknown FIT message field for a known message
// found in the official profile. It contains the message and field number, in
// addition to how many times the field for a specific message was encountered
// during decoding.
type UnknownField struct {
	MesgNum  MesgNum
	FieldNum byte
	Count    int
}

type unknownFieldSlice []UnknownField

func (p unknownFieldSlice) Len() int      { return len(p) }
func (p unknownFieldSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p unknownFieldSlice) Less(i, j int) bool {
	if p[i].MesgNum < p[j].MesgNum {
		return true
	} else if p[i].MesgNum > p[j].MesgNum {
		return false
	}
	return p[i].FieldNum < p[j].FieldNum
}

// UnknownMessage represents an unknown FIT message not found in the official
// profile. It contains the message number and how many times the was
// encountered during decoding.
type UnknownMessage struct {
	MesgNum MesgNum
	Count   int
}

type unknownMessageSlice []UnknownMessage

func (p unknownMessageSlice) Len() int      { return len(p) }
func (p unknownMessageSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p unknownMessageSlice) Less(i, j int) bool {
	return p[i].MesgNum < p[j].MesgNum
}
