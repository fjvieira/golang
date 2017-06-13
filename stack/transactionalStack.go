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

/*TransactionalStack A transational stack that can stack transactions contains operations (pop and push).
This stack can operate without any open transaction. You can open as many transactions as it is necessary.
In this case, it is a transaction inside another one and a Rollback and a commit will afect the imediately
underlaying transaction.
*/
type TransactionalStack struct {
	transactions *Stack
	stack        *Stack
}

type transaction struct {
	stack *Stack
}

type transactionItem struct {
	pushed bool
	value  interface{}
}

//CreateTransactionalStack Creates a new transaction stack.
func CreateTransactionalStack() *TransactionalStack {
	return &TransactionalStack{new(Stack), new(Stack)}
}

//OpenTransaction creates a new transaction.
func (t *TransactionalStack) OpenTransaction() {
	t.transactions.Push(transaction{new(Stack)})
}

//Commit Commits the last created transaction.
func (t *TransactionalStack) Commit() {
	currentTransaction, okCurrent := t.transactions.Pop().(transaction)

	if okCurrent {
		lastTransaction, okLast := t.transactions.Peek().(transaction)

		if okLast {
			for item := currentTransaction.stack.Pop(); item != nil; item = currentTransaction.stack.Pop() {
				lastTransaction.stack.Push(item)
			}
		}
	}
}

//Rollback Rollbacks - cancel - the last created transaction
func (t *TransactionalStack) Rollback() {
	currentTransaction, okCurrent := t.transactions.Pop().(transaction)

	if okCurrent {
		lastTransaction, okLast := t.transactions.Peek().(transaction)

		for item := currentTransaction.stack.Pop(); item != nil; item = currentTransaction.stack.Pop() {
			if okLast {
				lastTransaction.stack.Push(item)
			}
			transactionItem, ok := item.(transactionItem)
			if ok {
				if transactionItem.pushed {
					_ = t.stack.Pop()
				} else {
					t.stack.Push(transactionItem.value)
				}
			}
		}

	}
}

//Push Pushs a value from a top of the stack and include in the topmost transaction (if there is one opened)
func (t *TransactionalStack) Push(value interface{}) interface{} {
	currentTransaction, ok := t.transactions.Peek().(transaction)
	if ok {
		currentTransaction.stack.Push(transactionItem{true, value})
	}
	return t.stack.Push(value)
}

//Peek Returns the value in the top of the queue without remove it.
func (t *TransactionalStack) Peek() interface{} {
	return t.stack.Peek()
}

//Pop Removes the value in the top of stack and include in the topmost transaction (if there is one opened)
func (t *TransactionalStack) Pop() interface{} {
	currentTransaction, ok := t.transactions.Peek().(transaction)
	value := t.stack.Pop()
	if ok {
		currentTransaction.stack.Push(transactionItem{false, value})
	}
	return value
}
