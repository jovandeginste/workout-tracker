package convert

import (
	pb "github.com/ringsaturn/tzf/gen/go/tzf/v1"
	"github.com/tidwall/geojson/geometry"
)

func FromTimezonePBToGeometryPoly(timezone *pb.Timezone) []*geometry.Poly {
	ret := []*geometry.Poly{}
	for _, polygon := range timezone.Polygons {

		newPoints := make([]geometry.Point, 0)
		for _, point := range polygon.Points {
			newPoints = append(newPoints, geometry.Point{
				X: float64(point.Lng),
				Y: float64(point.Lat),
			})
		}

		holes := [][]geometry.Point{}
		for _, holePoly := range polygon.Holes {
			newHolePoints := make([]geometry.Point, 0)
			for _, point := range holePoly.Points {
				newHolePoints = append(newHolePoints, geometry.Point{
					X: float64(point.Lng),
					Y: float64(point.Lat),
				})
			}
			holes = append(holes, newHolePoints)
		}

		newPoly := geometry.NewPoly(newPoints, holes, nil)
		ret = append(ret, newPoly)
	}
	return ret
}
