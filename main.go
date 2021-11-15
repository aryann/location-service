package main

import (
	"fmt"

	"github.com/golang/geo/s2"
)

const (
	radiusMeters = 6371010.0
)

func main() {
	// For now, this is just a placeholder that computes the distance, in
	// meters, between two points in Seattle, WA.
	a := s2.PointFromLatLng(s2.LatLngFromDegrees(47.618698, -122.320229))
	b := s2.PointFromLatLng(s2.LatLngFromDegrees(47.615278, -122.320114))
	fmt.Println(a.Distance(b).Radians() * radiusMeters)
}
