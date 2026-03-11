package distance

import (
	"broken_sprockets/testHelpers"
	"testing"
)

func TestFeetToMeters(t *testing.T) {
	got := ToMeters(Feet(1))
	want := 0.3048

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMetersToFeet(t *testing.T) {
	got := ToFeet(Meters(1))
	want := 3.280839895013123

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMetersToKilometers(t *testing.T) {
	got := ToKilometers(Meters(1))
	want := 0.001

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMilesToMeters(t *testing.T) {
	got := ToMeters(Miles(1))
	want := 1609.344

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestDistanceAdd(t *testing.T) {
	got := ToMeters(Feet(10).Add(Meters(1)))
	want := 4.048

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
