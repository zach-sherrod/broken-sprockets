package frequency

import (
	"testing"

	"github.com/zach-sherrod/broken_sprockets/duration"
	"github.com/zach-sherrod/broken_sprockets/testHelpers"
)

func TestKilohertzToHertz(t *testing.T) {
	got := ToHertz(Kilohertz(1.0))
	want := 1000.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMegahertzToHertz(t *testing.T) {
	got := ToHertz(Megahertz(2.5))
	want := 2_500_000.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestRPMToHertz(t *testing.T) {
	got := ToHertz(RPM(120.0))
	want := 2.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestHertzToRPM(t *testing.T) {
	got := ToRPM(Hertz(2.0))
	want := 120.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestFromPeriod(t *testing.T) {
	got := ToHertz(FromPeriod(duration.Milliseconds(20)))
	want := 50.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestTryFromPeriodZeroDuration(t *testing.T) {
	_, err := TryFromPeriod(duration.Seconds(0.0))
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestToPeriod(t *testing.T) {
	got := duration.ToMilliseconds(ToPeriod(Hertz(100.0)))
	want := int64(10)

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestTryToPeriodZeroFrequency(t *testing.T) {
	_, err := TryToPeriod(Hertz(0.0))
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestFrequencyAdd(t *testing.T) {
	got := ToHertz(Hertz(10.0).Add(Kilohertz(1.0)))
	want := 1010.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
