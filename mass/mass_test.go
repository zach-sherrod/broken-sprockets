package mass

import (
	"testing"

	"github.com/zach-sherrod/broken_sprockets/testHelpers"
)

func TestPoundsToKilograms(t *testing.T) {
	got := ToKilograms(Pounds(1.0))
	want := 0.45359237

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestKilogramsToPounds(t *testing.T) {
	got := ToPounds(Kilograms(1.0))
	want := 2.2046226218487757

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestOuncesToKilograms(t *testing.T) {
	got := ToKilograms(Ounces(16.0))
	want := 0.45359237

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestGramsToKilograms(t *testing.T) {
	got := ToKilograms(Grams(1000.0))
	want := 1.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMassAdd(t *testing.T) {
	got := ToKilograms(Kilograms(2.0).Add(Pounds(10.0)))
	want := 2.0 + 4.5359237

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
