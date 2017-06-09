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

//Package stack Implements a simple stack.
package stack

import (
	"sync"
)

//Stack A simple stack implementation
type Stack struct {
	top    *item
	locker sync.Mutex
}

type item struct {
	previous *item
	value    interface{}
}

//Push Pushes a value to the top of the queue and return the previous top.
func (stack *Stack) Push(value interface{}) interface{} {
	stack.locker.Lock()
	defer stack.locker.Unlock()

	top := stack.top
	newItem := item{top, value}
	stack.top = &newItem
	if top == nil {
		return nil
	}
	return top.value
}

//Peek Returns the value in the top of the queue without remove it.
func (stack *Stack) Peek() interface{} {
	if stack.top == nil {
		return nil
	}
	return stack.top.value
}

//Pop Returns the value in the top of the queue and remove it.
func (stack *Stack) Pop() interface{} {
	stack.locker.Lock()
	defer stack.locker.Unlock()

	top := stack.top

	if top == nil {
		return nil
	}
	stack.top = top.previous
	return top.value
}
