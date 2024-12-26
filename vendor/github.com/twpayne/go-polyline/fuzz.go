// +build gofuzz

package polyline

func Fuzz(data []byte) int {
	_, rest, err := DecodeCoords(data)
	if err != nil || len(rest) > 0 {
		return 0
	}
	return 1
}
