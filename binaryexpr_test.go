package binaryexpr_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/Konstantin8105/binaryexpr"
)

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

func Test(t *testing.T) {
	t.Run("not filename", func(t *testing.T) {
		err := binaryexpr.Test("./some not exist file.go")
		if err == nil {
			t.Error("error not found")
		}
	})

	t.Run("not valid expressions", func(t *testing.T) {
		err := binaryexpr.Test("./binaryexpr_test.go")
		if err == nil {
			t.Error("error not found")
		}
	})

	t.Run("acceptable test", func(t *testing.T) {
		err := binaryexpr.Test("./binaryexpr.go")
		if err != nil {
			t.Error("error not found")
		}
	})
}

func Example() {
	err := binaryexpr.Test("./binaryexpr_test.go")
	fmt.Fprintf(os.Stdout, "%v", err)
	// Output:
	// ./binaryexpr_test.go
	// ├──./binaryexpr_test.go:13:5:	not acceprable operator: >=
	// ├──./binaryexpr_test.go:16:8:	not acceprable operator: >=
	// └──./binaryexpr_test.go:16:20:	not acceprable operator: >
}
