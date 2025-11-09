package geom

import (
	"fmt"
	"testing"
)

// Seems brittle :\
func TestGreatCircleDistance(t *testing.T) {
	// Test that SEA and SFO are ~ 1091km apart, accurate to 100 meters.
	sea := &LngLat{Latitude: 47.4489, Longitude: -122.3094}
	sfo := &LngLat{Latitude: 37.6160933, Longitude: -122.3924223}
	sfoToSea := 1093379.199082169

	dist := sea.GreatCircleDistance(sfo)

	if !(dist < (sfoToSea+0.1) && dist > (sfoToSea-0.1)) {
		t.Error("Unnacceptable result.", dist)
	}
}

func TestPointAtDistanceAndBearing(t *testing.T) {
	sea := &LngLat{Latitude: 47.44745785, Longitude: -122.308065668024}
	p := sea.PointAtDistanceAndBearing(1090700, 180)

	// Expected results of transposing point
	// ~1091km at bearing of 180 degrees
	resultLat := 37.638557
	resultLng := -122.308066

	withinLatBounds := p.Latitude < resultLat+0.001 && p.Latitude > resultLat-0.001
	withinLngBounds := p.Longitude < resultLng+0.001 && p.Longitude > resultLng-0.001
	if !(withinLatBounds && withinLngBounds) {
		t.Error("Unacceptable result.", fmt.Sprintf("[%f, %f]", p.Latitude, p.Longitude))
	}
}

func TestBearingTo(t *testing.T) {
	p1 := &LngLat{Latitude: 40.7486, Longitude: -73.9864}
	p2 := &LngLat{Latitude: 0.0, Longitude: 0.0}
	bearing := p1.BearingTo(p2)

	// Expected bearing 60 degrees
	resultBearing := 100.610833

	withinBearingBounds := bearing < resultBearing+0.001 && bearing > resultBearing-0.001
	if !withinBearingBounds {
		t.Error("Unacceptable result.", fmt.Sprintf("%f", bearing))
	}
}

func TestMidpointTo(t *testing.T) {
	p1 := &LngLat{Latitude: 52.205, Longitude: 0.119}
	p2 := &LngLat{Latitude: 48.857, Longitude: 2.351}

	p := p1.MidpointTo(p2)

	// Expected midpoint 50.5363°N, 001.2746°E
	resultLat := 50.53632
	resultLng := 1.274614

	withinLatBounds := p.Latitude < resultLat+0.001 && p.Latitude > resultLat-0.001
	withinLngBounds := p.Longitude < resultLng+0.001 && p.Longitude > resultLng-0.001
	if !(withinLatBounds && withinLngBounds) {
		t.Error("Unacceptable result.", fmt.Sprintf("[%f, %f]", p.Latitude, p.Longitude))
	}
}
