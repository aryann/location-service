package locationservice

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

// BenchmarkFindClosest benchmarks the FindClosest() in NaiveLocationService
// using 10,000 locations. NaiveLocationService performs a linear scan over the
// locations.
//
// On an Apple M1 computer, we expect results similar to:
//
//  BenchmarkFindClosest-8  4653    219324  ns/op   360 B/op    4 allocs/op
//
// In other words, the average FindClosest() operation took around 0.219
// milliseconds.
func BenchmarkFindClosest(b *testing.B) {
	f, err := os.Open("test-data.csv")
	if err != nil {
		b.Fatalf("could not open file: %v", err)
	}
	defer f.Close()

	service := NaiveLocationService{}

	reader := csv.NewReader(f)
	reader.Read() // Discard header.
	for {
		cols, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			b.Fatalf("could not read row: %v", err)
		}

		address := fmt.Sprintf("%s, %s, %s", cols[0], cols[1], cols[2])
		latitude, _ := strconv.ParseFloat(cols[4], 64)
		longitude, _ := strconv.ParseFloat(cols[5], 64)
		service.Add(address, GeoPoint{Latitude: latitude, Longitude: longitude})
	}

	maxDistanceMeters := 10000.0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		service.FindClosest(GeoPoint{47.618698, -122.320229}, maxDistanceMeters)
	}
}
