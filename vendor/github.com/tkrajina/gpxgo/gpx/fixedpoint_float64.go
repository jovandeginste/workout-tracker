package gpx

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// formattedFloat forces XML attributes to be marshalled as a fixed point decimal with 10 decimal places.
type formattedFloat float64

func (f formattedFloat) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  xml.Name{Local: name.Local},
		Value: strings.TrimRight(fmt.Sprintf("%.10f", f), "0."),
	}, nil
}
