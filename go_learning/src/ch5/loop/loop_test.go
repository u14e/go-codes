package loop

import (
	"testing"
	"time"
)

func TestWhileLoop(t *testing.T) {
	n := 0
	// while n < 5
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestInfiniteLoop(t *testing.T) {
	n := 0
	// while True
	for {
		t.Log(n)
		n++
		time.Sleep(2)
	}
}
