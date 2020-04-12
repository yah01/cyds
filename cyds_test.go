package cyds

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	fmt.Println("Test Queue:")
	var q Queue
	q.Push(5)
	q.Push("str")
	fmt.Println(q)
}
