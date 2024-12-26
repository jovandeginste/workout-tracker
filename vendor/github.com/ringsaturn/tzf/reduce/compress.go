package reduce

import (
	"fmt"

	"github.com/ringsaturn/tzf/pb"
	"github.com/twpayne/go-polyline"
)

func CompressedPointsToPolylineBytes(points []*pb.Point) []byte {
	expect := [][]float64{}
	for _, point := range points {
		expect = append(expect, []float64{float64(point.Lng), float64(point.Lat)})
	}
	return polyline.EncodeCoords(expect)
}

func DecompressedPolylineBytesToPoints(input []byte) []*pb.Point {
	expect := []*pb.Point{}
	coords, _, _ := polyline.DecodeCoords(input)
	for _, coord := range coords {
		expect = append(expect, &pb.Point{
			Lng: float32(coord[0]), Lat: float32(coord[1]),
		})
	}
	return expect
}

func CompressWithPolyline(input *pb.Timezones) *pb.CompressedTimezones {
	output := &pb.CompressedTimezones{
		Method:  pb.CompressMethod_Polyline,
		Version: input.Version,
	}
	for _, timezone := range input.Timezones {
		reducedTimezone := &pb.CompressedTimezone{
			Name: timezone.Name,
		}
		for _, polygon := range timezone.Polygons {
			newPoly := &pb.CompressedPolygon{
				Points: CompressedPointsToPolylineBytes(polygon.Points),
				Holes:  make([]*pb.CompressedPolygon, 0),
			}
			for _, hole := range polygon.Holes {
				newPoly.Holes = append(newPoly.Holes, &pb.CompressedPolygon{
					Points: CompressedPointsToPolylineBytes(hole.Points),
				})
			}
			reducedTimezone.Data = append(reducedTimezone.Data, newPoly)
		}
		output.Timezones = append(output.Timezones, reducedTimezone)
	}
	return output
}

func Compress(input *pb.Timezones, method pb.CompressMethod) (*pb.CompressedTimezones, error) {
	switch method {
	case pb.CompressMethod_Polyline:
		return CompressWithPolyline(input), nil
	default:
		return nil, fmt.Errorf("tzf/reduce: unknown method %v", method)
	}
}

func DecompressWithPolyline(input *pb.CompressedTimezones) *pb.Timezones {
	output := &pb.Timezones{
		Version: input.Version,
	}
	for _, timezone := range input.Timezones {
		reducedTimezone := &pb.Timezone{
			Name: timezone.Name,
		}
		for _, polygon := range timezone.Data {
			newPoly := &pb.Polygon{
				Points: DecompressedPolylineBytesToPoints(polygon.Points),
				Holes:  make([]*pb.Polygon, 0),
			}
			for _, hole := range polygon.Holes {
				newPoly.Holes = append(newPoly.Holes, &pb.Polygon{
					Points: DecompressedPolylineBytesToPoints(hole.Points),
				})
			}
			reducedTimezone.Polygons = append(reducedTimezone.Polygons, newPoly)
		}
		output.Timezones = append(output.Timezones, reducedTimezone)
	}
	return output
}

func Decompress(input *pb.CompressedTimezones) (*pb.Timezones, error) {
	switch input.Method {
	case pb.CompressMethod_Polyline:
		return DecompressWithPolyline(input), nil
	default:
		return nil, fmt.Errorf("tzf/reduce: unknown method %v", input.Method)
	}
}
