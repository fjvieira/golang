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

import (
	"testing"
)

func TestCommitWithoutOpenTrans(t *testing.T) {
	stack := CreateTransactionalStack()

	stack.Commit()

	stack.Commit()
}

func TestCommitOpenEmptyTrans(t *testing.T) {
	stack := CreateTransactionalStack()

	stack.OpenTransaction()

	stack.Commit()

	if stack.transactions.Peek() != nil {
		t.Error("Something wrong in the transactions control")
	}

}

func TestRollbackWithoutOpenTrans(t *testing.T) {
	stack := CreateTransactionalStack()

	stack.Rollback()

	stack.Rollback()

	if stack.Peek() != nil {
		t.Error("Something wrong in the transactions control")
	}
}

func TestRollbackOpenEmptyTrans(t *testing.T) {
	stack := CreateTransactionalStack()

	stack.OpenTransaction()

	stack.Rollback()

	if stack.transactions.Peek() != nil {
		t.Error("Something wrong in the transactions control")
	}
}

func TestCommitRollbackTrans(t *testing.T) {
	stack := CreateTransactionalStack()

	stack.OpenTransaction()

	stack.Push(0)

	stack.Push(1)

	stack.OpenTransaction()

	stack.Push(2)

	stack.Rollback()

	_ = stack.Pop()

	if stack.Peek() != 0 {
		t.Error("Something wrong in the transactions control - not 0 after pop")
	}

	stack.Push(3)

	stack.Rollback()

	if stack.Peek() != nil {
		t.Error("Something wrong in the transactions control - not nil")
	}

}

func TestMoreCommitRollbackTrans(t *testing.T) {
	stack := CreateTransactionalStack()

	stack.OpenTransaction()

	stack.Push(0)

	stack.Commit()

	stack.Rollback()

	if stack.Peek() != 0 {
		t.Error("Something wrong in the transactions control - dont kept the item ")
	}

	_ = stack.Pop()

}

func TestChainedCommitRollbackTrans(t *testing.T) {
	stack := CreateTransactionalStack()

	stack.OpenTransaction()

	stack.Push(0)

	stack.Commit()

	stack.Rollback()

	if stack.Peek() != 0 {
		t.Error("Something wrong in the transactions control - dont kept the item ")
	}

	_ = stack.Pop()

	stack.OpenTransaction()

	stack.Push(0)

	stack.Push(1)

	stack.OpenTransaction()

	stack.Push(2)

	stack.Commit()

	stack.Rollback()

	if stack.Peek() != nil {
		t.Error("Something wrong in the transactions control - should not have items")
	}

}
