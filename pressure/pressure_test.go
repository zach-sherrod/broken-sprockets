package pressure

import (
	"testing"

	"github.com/zach-sherrod/broken-sprockets/testHelpers"
)

func TestKilopascalsToPascals(t *testing.T) {
	got := ToPascals(Kilopascals(1.0))
	want := 1000.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestBarToPascals(t *testing.T) {
	got := ToPascals(Bar(1.0))
	want := 100000.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestPSIToPascals(t *testing.T) {
	got := ToPascals(PSI(1.0))
	want := 6894.757293168

	if !testHelpers.AlmostEqual(got, want, 1e-9) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestAtmospheresToPascals(t *testing.T) {
	got := ToPascals(Atmospheres(1.0))
	want := 101325.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestTorrToPascals(t *testing.T) {
	got := ToPascals(Torr(760.0))
	want := 101325.0

	if !testHelpers.AlmostEqual(got, want, 1e-6) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestPascalsToPSI(t *testing.T) {
	got := ToPSI(Pascals(6894.757293168))
	want := 1.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestPressureAdd(t *testing.T) {
	got := ToPascals(Pascals(500).Add(Kilopascals(1.5)))
	want := 2000.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
