package locationservice

import "github.com/golang/geo/s2"

type pointContainer struct {
	id    string
	point s2.Point
}

type NaiveLocationService struct {
	locations []pointContainer
}

func (n *NaiveLocationService) Add(id string, point GeoPoint) {
	n.locations = append(n.locations, pointContainer{id, toS2Point(point)})
}

func (n *NaiveLocationService) FindClosest(point GeoPoint, maxDistance float64) []SearchResult {
	target := toS2Point(point)

	result := []SearchResult{}
	for _, location := range n.locations {
		distance := target.Distance(location.point).Radians() * earthRadiusMeters
		if distance <= maxDistance {
			result = append(result, SearchResult{location.id, int(distance)})
		}
	}
	return result
}
