package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("T", 10)
	expected := "TTTTTTTTTT"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat(t *testing.T) {
	repeated := Repeat("go", 3)
	fmt.Println(repeated)
	// Output:
	// gogogo
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 20)
	}
}
