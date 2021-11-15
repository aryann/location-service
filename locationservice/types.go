package locationservice

import "github.com/golang/geo/s2"

const (
	earthRadiusMeters = 6371010.0
)

type SearchResult struct {
	ID             string
	DistanceMeters int
}

// GeoPoint is a type for representing a latitude and longitude in decimal form.
type GeoPoint struct {
	Latitude, Longitude float64
}

// LocationService is a type for storing a set of locations on Earth and
// searching for them based off of a GeoPoint and a maximum distance.
type LocationService interface {
	Add(id string, point GeoPoint)
	FindClosest(point GeoPoint, maxDistanceMeters int) []SearchResult
}

func toS2Point(point GeoPoint) s2.Point {
	return s2.PointFromLatLng(s2.LatLngFromDegrees(point.Latitude, point.Longitude))
}
