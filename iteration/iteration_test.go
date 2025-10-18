package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expect := "aaaaaa"

	if repeated != expect {
		t.Errorf("expected %q but got %q", expect, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// setup code can go here
	for b.Loop() {
		// code to test
		Repeat("a", 6)
	}
	// cleanup code can go here
}
