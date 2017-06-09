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

//Package simpleThreadPool implements a simple thread pool that receives tasks implementing Runnable interface.
package simpleThreadPool

import (
	"errors"
	"sync/atomic"
)

var (
	errInactive      = errors.New("The pool is not active")
	errAlreadyActive = errors.New("The pool has already been active")
)

//Runnable A Java-like interface to be used in the task queue
type Runnable interface {
	Run()
}

//A representation of a worker.
type worker struct {
	inputChan chan Runnable
}

//Activates a worker.
func (worker *worker) activate() {
	worker.inputChan = make(chan Runnable)
	go worker.run()
}

func (worker *worker) run() {
	for {
		job, ok := <-worker.inputChan
		if ok {
			job.Run()
		} else {
			return
		}
	}
}

//Deactivates a worker.
func (worker *worker) desactivate() {
	close(worker.inputChan)
}

//ThreadPool A simple thread pool that run tasks implementing the Runnable interface.
type ThreadPool struct {
	nActiveWorkers uint32
	bufferSize     int
	active         uint32
	currentWorker  uint32
	workers        []worker
	jobChan        chan Runnable
}

//Activate Activates the thread pool.
func (pool *ThreadPool) Activate() error {
	if pool.active == 0 {
		pool.jobChan = make(chan Runnable, pool.bufferSize)
		pool.workers = make([]worker, 0)
		for i := uint32(0); i < pool.nActiveWorkers; i++ {
			w := new(worker)
			w.activate()
			pool.workers = append(pool.workers, *w)
			go w.run()
		}
		atomic.SwapUint32(&pool.active, 1)
		go pool.run()
		return nil
	}
	return errAlreadyActive
}

func (pool *ThreadPool) run() {
	for {
		job, ok := <-pool.jobChan
		if ok {
			atomic.AddUint32(&pool.currentWorker, 1)
			atomic.CompareAndSwapUint32(&pool.currentWorker, pool.nActiveWorkers, 0)
			pool.workers[pool.currentWorker].inputChan <- job
		} else {
			return
		}
	}
}

//Deactivate Deactivates the thread pool.
func (pool *ThreadPool) Deactivate() error {
	if pool.active == 1 {
		close(pool.jobChan)
		for _, w := range pool.workers {
			w.desactivate()
		}
		atomic.SwapUint32(&pool.active, 0)
		return nil
	}
	return errInactive
}

//Execute Executes a task through the pool.
func (pool *ThreadPool) Execute(task Runnable) {
	pool.jobChan <- task
}
