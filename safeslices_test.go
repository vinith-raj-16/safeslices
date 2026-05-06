package safeslices

import "testing"

func TestMaxSafe(t *testing.T) {
    max, ok := MaxSafe([]int{1, 5, 3})

    if !ok {
        t.Fatal("expected ok=true")
    }

    if max != 5 {
        t.Fatal("expected max=5")
    }
}