package cyDS

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	fmt.Println("Test Queue:")
	var q Queue
	fmt.Println("Push", 5)
	q.Push(5)
	v := q.Front()
	fmt.Println(v)
}
