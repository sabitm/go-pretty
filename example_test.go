package pretty_test

import (
	"fmt"

	"github.com/sabitm/go-pretty"
)

func Example() {
	//nolint:structcheck
	type myType struct {
		a, b int
	}
	x := []myType{{1, 2}, {3, 4}, {5, 6}}
	fmt.Printf("%# v", pretty.Formatter(x))
	// output:
	// []pretty_test.myType{
	//     {a:1, b:2},
	//     {a:3, b:4},
	//     {a:5, b:6},
	// }
}
