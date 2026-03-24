package duration

import (
	"testing"

	"github.com/zach-sherrod/broken_sprockets/testHelpers"
)

func TestMillisecondsToSeconds(t *testing.T) {
	got := ToSeconds(Milliseconds(1500))
	want := 1.5

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestSecondsToMilliseconds(t *testing.T) {
	got := ToMilliseconds(Seconds(2.5))
	want := int64(2500)

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMinutesToSeconds(t *testing.T) {
	got := ToSeconds(Minutes(2))
	want := 120.0

	if !testHelpers.AlmostEqual(got, want, 1e-12) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestDurationAdd(t *testing.T) {
	got := ToMilliseconds(Seconds(1).Add(Milliseconds(250)))
	want := int64(1250)

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
