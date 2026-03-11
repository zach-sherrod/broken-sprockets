package angularvelocity

import (
	"broken_sprockets/angle"
	"broken_sprockets/duration"
	"broken_sprockets/testHelpers"
	"math"
	"testing"
)

func TestDegreesPerSecondToRadiansPerSecond(t *testing.T) {
	got := ToRadiansPerSecond(DegreesPerSecond(180.0))
	want := math.Pi

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestRadiansPerSecondToDegreesPerSecond(t *testing.T) {
	got := ToDegreesPerSecond(RadiansPerSecond(math.Pi / 2.0))
	want := 90.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestArcSecondsPerSecondToDegreesPerSecond(t *testing.T) {
	got := ToDegreesPerSecond(ArcSecondsPerSecond(3600.0))
	want := 1.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestDegreesPerSecondToArcSecondsPerSecond(t *testing.T) {
	got := ToArcSecondsPerSecond(DegreesPerSecond(1.0))
	want := 3600.0

	if !testHelpers.AlmostEqual(got, want, 1e-9) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestFromAngleOverDuration(t *testing.T) {
	got := ToDegreesPerSecond(
		FromAngleOverDuration(angle.Degrees(90.0), duration.Seconds(2.0)),
	)
	want := 45.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestTryFromAngleOverDurationZeroDuration(t *testing.T) {
	_, err := TryFromAngleOverDuration(angle.Degrees(10.0), duration.Seconds(0.0))
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestToAngleOverDuration(t *testing.T) {
	got := angle.ToDegrees(
		ToAngleOverDuration(DegreesPerSecond(30.0), duration.Seconds(3.0)),
	)
	want := 90.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestAngularVelocityAdd(t *testing.T) {
	got := ToDegreesPerSecond(DegreesPerSecond(10.0).Add(ArcMinutesPerSecond(30.0)))
	want := 10.5

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
