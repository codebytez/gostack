// Package gostack provides simple stack implementation
package gostack

import (
	"errors"
	"reflect"
)

type Stack struct {
	items []interface{}
}

// NewStack creates a new Stack
func NewStack() *Stack {
	st := new(Stack)
	return st
}

// Push pushes an element to the top of the stack
func (st *Stack) Push(val interface{}) {
	st.items = append(st.items, val)
}

// Pop pops the last element in the stack
func (st *Stack) Pop() (interface{}, error) {
	if len(st.items) > 0 {
		val := st.items[len(st.items)-1]
		st.items = st.items[0:(len(st.items) - 1)]
		return val, nil
	}

	return nil, errors.New("Empty stack")
}

// Peek returns the top item in the stack without removing it
func (st *Stack) Peek() (interface{}, error) {
	if len(st.items) > 0 {
		return st.items[len(st.items)-1], nil
	}

	return nil, errors.New("Empty stack")
}

// Contains checks if the input value exists in the stack
func (st *Stack) Contains(checkVal interface{}) bool {
	for _, val := range st.items {
		if reflect.DeepEqual(checkVal, val) {
			return true
		}
	}

	return false
}

// Synchronized returns true if the stack is thread safe
func (st *Stack) Synchronized() bool {
	return false
}

// Clear clears all elements from the stack
func (st *Stack) Clear() {
	st.items = st.items[:0]
}

// Counts returns the number of elements in the stack
func (st *Stack) Count() int {
	return len(st.items)
}
