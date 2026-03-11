package acceleration

import (
	"broken_sprockets/duration"
	"broken_sprockets/speed"
	"broken_sprockets/testHelpers"
	"testing"
)

func TestFeetPerSecondSquaredToMetersPerSecondSquared(t *testing.T) {
	got := ToMetersPerSecondSquared(FeetPerSecondSquared(1.0))
	want := 0.3048

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMetersPerSecondSquaredToFeetPerSecondSquared(t *testing.T) {
	got := ToFeetPerSecondSquared(MetersPerSecondSquared(1.0))
	want := 3.280839895013123

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestGToMetersPerSecondSquared(t *testing.T) {
	got := ToMetersPerSecondSquared(G(1.0))
	want := 9.80665

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMetersPerSecondSquaredToG(t *testing.T) {
	got := ToG(MetersPerSecondSquared(9.80665))
	want := 1.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestFromSpeedOverDuration(t *testing.T) {
	got := ToMetersPerSecondSquared(
		FromSpeedOverDuration(speed.MetersPerSecond(20.0), duration.Seconds(4.0)),
	)
	want := 5.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestTryFromSpeedOverDurationZeroDuration(t *testing.T) {
	_, err := TryFromSpeedOverDuration(speed.MetersPerSecond(10.0), duration.Seconds(0.0))
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestToSpeedOverDuration(t *testing.T) {
	got := speed.ToMetersPerSecond(
		ToSpeedOverDuration(MetersPerSecondSquared(3.0), duration.Seconds(5.0)),
	)
	want := 15.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestAccelerationAdd(t *testing.T) {
	got := ToMetersPerSecondSquared(MetersPerSecondSquared(2.0).Add(FeetPerSecondSquared(10.0)))
	want := 2.0 + 3.048

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
