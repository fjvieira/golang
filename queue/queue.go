/*
ISC License

Copyright (c) 2017, Fernando Jose Vieira

Permission to use, copy, modify, and/or distribute this software for any
purpose with or without fee is hereby granted, provided that the above
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE
OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
PERFORMANCE OF THIS SOFTWARE.
*/

//Package queue implements a simple thread safe queue.
package queue

import "sync"

type item struct {
	next  *item
	value interface{}
}

//Queue A thread-safe queue
type Queue struct {
	head   *item
	last   *item
	locker sync.Mutex
}

// Add Adds a Runnable at the end of the queue.
func (queue *Queue) Add(value interface{}) {
	queue.locker.Lock()
	defer queue.locker.Unlock()

	newItem := &item{nil, value}
	if queue.head == nil {
		queue.head = newItem
	}
	if queue.last != nil {
		queue.last.next = newItem
	}
	queue.last = newItem
}

//Remove Removes the oldest item of the queue.
func (queue *Queue) Remove() interface{} {
	queue.locker.Lock()
	defer queue.locker.Unlock()

	current := queue.head
	if current == nil {
		queue.last = nil
		return nil
	}
	queue.head = current.next
	return current.value
}

//Peek Return the item on the top of the queue without removing it.
func (queue *Queue) Peek() interface{} {
	return queue.head.value
}
