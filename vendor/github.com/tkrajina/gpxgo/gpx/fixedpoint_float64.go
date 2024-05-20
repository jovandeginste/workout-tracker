package gpx

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// formattedFloat forces XML attributes to be marshalled as a fixed point decimal with 10 decimal places.
type formattedFloat float64

func (f formattedFloat) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	s := strings.TrimRight(strconv.FormatFloat(float64(f), 'f', 10, 64), "0")
	if strings.HasSuffix(s, ".") {
		s += "0"
	}
	return xml.Attr{
		Name:  xml.Name{Local: name.Local},
		Value: s,
	}, nil
}
