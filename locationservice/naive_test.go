package locationservice

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func newService() (*NaiveLocationService, error) {
	f, err := os.Open("test-data.csv")
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer f.Close()

	service := &NaiveLocationService{}

	reader := csv.NewReader(f)
	reader.Read() // Discard header.
	for {
		cols, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("could not read row: %w", err)
		}

		address := fmt.Sprintf("%s, %s, %s", cols[0], cols[1], cols[2])
		latitude, _ := strconv.ParseFloat(cols[4], 64)
		longitude, _ := strconv.ParseFloat(cols[5], 64)
		service.Add(address, GeoPoint{Latitude: latitude, Longitude: longitude})
	}
	return service, nil
}

func TestFindClosest(t *testing.T) {
	service, err := newService()
	if err != nil {
		t.Fatalf("could not construct service: %v", err)
	}

	maxDistanceMeters := 10000
	result := service.FindClosest(GeoPoint{47.618698, -122.320229}, maxDistanceMeters)
	for _, location := range result {
		if !strings.Contains(location.ID, "Seattle") {
			t.Errorf("unexpected result: %+v", location)
		}
		if location.DistanceMeters > maxDistanceMeters {
			t.Errorf("unexpected distance: %+v", location)
		}
	}
}

// BenchmarkFindClosest benchmarks the FindClosest() function in
// NaiveLocationService using 10,000 locations. NaiveLocationService performs a
// linear scan over the locations.
//
// On an Apple M1 computer, we expect results similar to:
//
//  BenchmarkFindClosest-8  4653    219324  ns/op   360 B/op    4 allocs/op
//
// In other words, the average FindClosest() operation took around 0.219
// milliseconds.
//
// Note that the implementation represents points as a 3D vector. As a result,
// the distance calculation operations are significantly faster compared to a
// representation that retains the original latitude and longitude. We expect an
// implementation that does not use 3D vectors to be about three times slower:
//
//   BenchmarkFindClosest-8	1800	644824 ns/op	360 B/op	4 allocs/op
//
func BenchmarkFindClosest(b *testing.B) {
	service, err := newService()
	if err != nil {
		b.Fatalf("could not construct service: %v", err)
	}

	maxDistanceMeters := 10000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.FindClosest(GeoPoint{47.618698, -122.320229}, maxDistanceMeters)
	}
}
