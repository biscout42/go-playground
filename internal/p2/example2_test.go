package p2

import (
	"fmt"
	"playground/example/mutex/internal/util"
	"testing"
)

func TestExample2(t *testing.T) {
	fmt.Println("Testing example 2")
	util.Foo("example2")
	fmt.Println("Testing example 2 Done")
}

func TestExample21(t *testing.T) {
	fmt.Println("Testing example 21")
	util.Foo("example21")
	fmt.Println("Testing example 21 Done")
}
