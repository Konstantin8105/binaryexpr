package binaryexpr_test

import (
	"testing"

	"github.com/Konstantin8105/binaryexpr"
)

func Test(t *testing.T) {
	t.Run("not filename", func(t *testing.T) {
		tf := &testing.T{}
		binaryexpr.Test(tf, "./some not exist file.go")
	})

	t.Run("not valid expressions", func(t *testing.T) {
		tf := &testing.T{}
		binaryexpr.Test(tf, "./binaryexpr_test.go")
		if !tf.Failed() {
			t.Fatalf("errors are not found")
		}
	})
}

func testfunction() {
	a := 1
	if a >= 23 {
		a = 2
	}
	b := (a >= 77) || a > 5
	if b == true {
		b = a <= 33
	}
	_ = b
}
