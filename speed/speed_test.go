package speed

import (
	"testing"

	"github.com/zach-sherrod/broken_sprockets/distance"
	"github.com/zach-sherrod/broken_sprockets/duration"
	"github.com/zach-sherrod/broken_sprockets/testHelpers"
)

func TestMilesPerHourToMetersPerSecond(t *testing.T) {
	got := ToMetersPerSecond(MilesPerHour(1))
	want := 0.44704

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMetersPerSecondToMilesPerHour(t *testing.T) {
	got := ToMilesPerHour(MetersPerSecond(1))
	want := 2.2369362920544

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestKnotsToMetersPerSecond(t *testing.T) {
	got := ToMetersPerSecond(Knots(1))
	want := 1852.0 / 3600.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestFromDistanceOverDuration(t *testing.T) {
	got := ToMetersPerSecond(
		FromDistanceOverDuration(distance.Meters(100), duration.Seconds(10)),
	)
	want := 10.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestToDistanceOverDuration(t *testing.T) {
	got := distance.ToMeters(
		ToDistanceOverDuration(MetersPerSecond(10), duration.Seconds(5)),
	)
	want := 50.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestSpeedAdd(t *testing.T) {
	got := ToMetersPerSecond(MetersPerSecond(5).Add(FeetPerSecond(10)))
	want := 5.0 + 3.048

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
