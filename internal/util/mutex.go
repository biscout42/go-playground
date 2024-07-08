package util

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var count = 0

func Foo(name string) {
	fmt.Printf("### Mutex lock starts. count=[%d] %s \n", count, name)
	mu.Lock()
	defer func(m *sync.Mutex) {
		fmt.Printf("### Mutex Unlock starts. count=[%d] %s \n", count, name)
		m.Unlock()
	}(&mu)
	count = count + 1
	time.Sleep(5 * time.Second)
	fmt.Printf("### Mutex locked. count=[%d] %s \n", count, name)
}
