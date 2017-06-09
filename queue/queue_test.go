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

import (
	"testing"
	"time"
)

type Message struct {
	message string
}

func (task Message) Run() interface{} {
	return "test"
}

func TestCreateAndUseQueue(t *testing.T) {
	queue := Queue{}

	queue.Add(Message{"task1"})
	queue.Add(Message{"task2"})
	queue.Add(Message{"task3"})

	nextMessage := queue.Remove()
	testMessage(t, nextMessage, "task1")

	queue.Add(Message{"task4"})

	queue.Add(Message{"task5"})

	nextMessage = queue.Remove()
	testMessage(t, nextMessage, "task2")

	nextMessage = queue.Remove()
	testMessage(t, nextMessage, "task3")

	nextMessage = queue.Remove()
	testMessage(t, nextMessage, "task4")

	nextMessage = queue.Peek()
	testMessage(t, nextMessage, "task5")

	nextMessage = queue.Remove()
	testMessage(t, nextMessage, "task5")

	nextMessage = queue.Remove()
	if nextMessage != nil {
		t.Error("The queue should be empty")
	}
	nextMessage = queue.Remove()
	if nextMessage != nil {
		t.Error("The queue should be empty")
	}
}

func testMessage(t *testing.T, task interface{}, message string) {
	if task == nil {
		t.Error("Item was not added")
	}

	if task.(Message).message != message {
		t.Errorf("Item was not in the right order")
	}

}

func TestConcorrency(t *testing.T) {
	queue := Queue{}

	go addInQueue(&queue, "task1")
	go addInQueue(&queue, "task2")
	go addInQueue(&queue, "task3")
	go addInQueue(&queue, "task4")
	go addInQueue(&queue, "task5")

	time.Sleep(time.Millisecond)

	n := 0
	for queue.Remove() != nil {
		n++
	}
	if n != 5 {
		t.Errorf("Concurrency lock failed")
	}
}

func addInQueue(queue *Queue, s string) {
	queue.Add(s)
}
