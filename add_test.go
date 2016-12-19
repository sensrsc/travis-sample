package math

import "testing"

func TestAdd(t *testing.T) {
  const int1, int2, out = 1, 1, 2
  if x := Add(int1, int2); x != out {
    t.Errorf("Add(%v, %v) = %v, want %v", int1, int2, x, out)
  }
}
