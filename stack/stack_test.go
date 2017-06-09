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

package stack

import "testing"

func TestCreateStack(t *testing.T) {
	stack := new(Stack)

	value := stack.Peek()

	if value != nil {
		t.Error("The stack should be empty")
	}

	value = stack.Pop()

	if value != nil {
		t.Error("The stack should be empty")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if !(stack.Peek() == 3 && stack.Pop() == 3 && stack.Peek() == 2) {
		t.Error("The pop and push mech is not working")
	}

	stack.Push(4)

	if !(stack.Pop() == 4 && stack.Pop() == 2 && stack.Peek() == 1) {
		t.Error("The pop and push mech is not working")
	}

	if !(stack.Pop() == 1 && stack.Peek() == nil) {
		t.Error("The pop and push mech is not working")
	}

	if !(stack.Pop() == nil && stack.Peek() == nil) {
		t.Error("The pop and push mech is not working")
	}
}
