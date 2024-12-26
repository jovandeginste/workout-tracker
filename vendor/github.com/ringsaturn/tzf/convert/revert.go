package convert

import "github.com/ringsaturn/tzf/pb"

func FromPbPolygonToGeoMultipolygon(pbpoly []*pb.Polygon) MultiPolygonCoordinates {
	res := MultiPolygonCoordinates{}
	for _, poly := range pbpoly {
		newGeoPoly := make(PolygonCoordinates, 0)

		mainpoly := [][2]float64{}
		for _, point := range poly.Points {
			mainpoly = append(mainpoly, [2]float64{float64(point.Lng), float64(point.Lat)})
		}
		newGeoPoly = append(newGeoPoly, mainpoly)

		for _, holepoly := range poly.Holes {
			holepolyCoords := [][2]float64{}
			for _, point := range holepoly.Points {
				holepolyCoords = append(holepolyCoords, [2]float64{float64(point.Lng), float64(point.Lat)})
			}
			newGeoPoly = append(newGeoPoly, holepolyCoords)
		}
		res = append(res, newGeoPoly)
	}
	return res
}

func RevertItem(input *pb.Timezone) *FeatureItem {
	return &FeatureItem{
		Type: FeatureType,
		Properties: PropertiesDefine{
			Tzid: input.Name,
		},
		Geometry: GeometryDefine{
			Type:        MultiPolygonType,
			Coordinates: FromPbPolygonToGeoMultipolygon(input.Polygons),
		},
	}
}

// Revert could convert pb define data to GeoJSON format.
func Revert(input *pb.Timezones) *BoundaryFile {
	output := &BoundaryFile{}
	for _, timezone := range input.Timezones {
		item := RevertItem(timezone)
		output.Features = append(output.Features, item)
	}
	output.Type = "FeatureCollection"
	return output
}
