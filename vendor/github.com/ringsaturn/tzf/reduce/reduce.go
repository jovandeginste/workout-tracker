// Package reduce could reduce Polygon size both polygon lines and float precise.
package reduce

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/simplify"
	pb "github.com/ringsaturn/tzf/gen/go/tzf/v1"
)

func ReducePoints(points []*pb.Point) []*pb.Point {
	if len(points) == 0 {
		return points
	}
	original := orb.LineString{}
	for _, point := range points {
		original = append(original, orb.Point{float64(point.Lng), float64(point.Lat)})
	}
	reduced := simplify.DouglasPeucker(0.001).Simplify(original.Clone()).(orb.LineString)
	res := make([]*pb.Point, 0)
	for _, orbPoint := range reduced {
		res = append(res, &pb.Point{
			Lng: float32(orbPoint.Lon()),
			Lat: float32(orbPoint.Lat()),
		})
	}
	return res
}

func Do(input *pb.Timezones, skip int, precise float64, minist float64) *pb.Timezones {
	output := &pb.Timezones{
		Version: input.Version,
	}
	for _, timezone := range input.Timezones {
		reducedTimezone := &pb.Timezone{
			Name: timezone.Name,
		}
		for _, polygon := range timezone.Polygons {
			newPoly := &pb.Polygon{
				Points: ReducePoints(polygon.Points),
				Holes:  make([]*pb.Polygon, 0),
			}
			for _, hole := range polygon.Holes {
				newPoly.Holes = append(newPoly.Holes, &pb.Polygon{
					Points: ReducePoints(hole.Points),
				})
			}
			reducedTimezone.Polygons = append(reducedTimezone.Polygons, newPoly)
		}
		output.Timezones = append(output.Timezones, reducedTimezone)
	}
	return output
}
