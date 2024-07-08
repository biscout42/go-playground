package p1

import (
	"fmt"
	"playground/example/mutex/internal/util"
	"testing"
)

func TestExample1(t *testing.T) {
	fmt.Println("Testing example 1")
	util.Foo("example1")
	fmt.Println("Testing example 1 Done")
}
