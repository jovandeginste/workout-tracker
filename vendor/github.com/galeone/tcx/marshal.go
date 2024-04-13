package tcx

import (
	"encoding/xml"
)

func ToBytes(file TCXDB) ([]byte, error) {
	return xml.MarshalIndent(file, "", "  ")
}
