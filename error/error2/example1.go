// What does this program print?
package main

import (
	"errors"
	"fmt"
)

// A naive circular buffer
type CircularBuffer struct {
	buf             []int
	cap, head, tail int
}

// add adds the number to the tail of the buffer
func (q *CircularBuffer) add(d int) error {
	if q.tail-q.head > q.cap {
		return errors.New("buffer full")
	}
	q.buf[q.tail%q.cap] = d
	q.tail += 1
	return nil
}

// remove removes the number that is in the buffer the longest
func (q CircularBuffer) remove() (int, error) {
	if q.tail == q.head {
		return -1, errors.New("buffer empty")
	}
	x := q.buf[q.head%q.cap]
	q.head += 1
	return x, nil
}
func main() {
	q := CircularBuffer{buf: make([]int, 5), cap: 5}
	_ = q.add(1)
	act, _ := q.remove()
	fmt.Printf("removed: %d ", act)
	_ = q.add(2)
	act, _ = q.remove()
	fmt.Printf("removed: %d\n", act)
}
