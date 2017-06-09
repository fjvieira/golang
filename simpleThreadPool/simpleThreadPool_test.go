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
package simpleThreadPool

import (
	"fmt"
	"testing"
	"time"
)

var testWorker int

type Task struct {
	message string
}

func (task Task) Run() {
	testWorker++
}

func TestWorker(t *testing.T) {
	w := worker{}

	w.activate()

	for index := 0; index < 100; index++ {
		w.inputChan <- Task{fmt.Sprintf("test%d", index)}
	}

	time.Sleep(10 * time.Millisecond)

	w.desactivate()

	if testWorker != 100 {
		t.Error("Worker did not processed all the tasks")
	}
}

func TestThreadPool(t *testing.T) {
	pool := ThreadPool{}
	pool.bufferSize = 2
	pool.nActiveWorkers = 4

	pool.Activate()

	for index := 0; index < 100; index++ {
		pool.Execute(Task{fmt.Sprintf("test%d", index)})
	}

	time.Sleep(10 * time.Millisecond)

	pool.Deactivate()

	pool.Activate()

	for index := 0; index < 100; index++ {
		pool.Execute(Task{fmt.Sprintf("test%d", index)})
	}

	time.Sleep(10 * time.Millisecond)

	pool.Deactivate()
}
