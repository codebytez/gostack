package gostack

import "testing"

func TestPushPop(t *testing.T) {
	st := NewStack()

	val, err := st.Pop()

	if err == nil || val != nil {
		t.Error("Pop on empty stack should return error")
	}

	st.Push(1)
	st.Push("test")
	st.Push(2.0)

	val, err = st.Pop()
	if err != nil || val.(float64) != 2.0 {
		t.Errorf("PushPop Got - %f,expecting - 2.0", val.(float64))
	}

	val, err = st.Pop()
	if err != nil || val.(string) != "test" {
		t.Errorf("PushPop Got - %s,expecting - 'test'", val.(string))
	}

	val, err = st.Pop()
	if err != nil || val.(int) != 1 {
		t.Errorf("PushPop Got - %d,expecting - '1'", val.(int))
	}

	val, err = st.Pop()
	if err == nil || val != nil {
		t.Error("Pop on empty stack should return error")
	}

}

func TestPeek(t *testing.T) {
	st := NewStack()

	val, err := st.Peek()

	if err == nil || val != nil {
		t.Error("Peek on empty stack should return error")
	}

	st.Push(1)
	st.Push("test")

	val, err = st.Peek()
	if err != nil || val.(string) != "test" {
		t.Errorf("PushPop Got - %s,expecting - 'test'", val.(string))
	}

	st.Pop()
	val, err = st.Peek()
	if err != nil || val.(int) != 1 {
		t.Errorf("PushPop Got - %d,expecting - '1'", val.(int))
	}

	st.Push(2.0)
	val, err = st.Peek()
	if err != nil || val.(float64) != 2.0 {
		t.Errorf("PushPop Got - %f,expecting - 2.0", val.(float64))
	}

}

func TestContains(t *testing.T) {
	st := NewStack()

	yes := st.Contains("test")
	if yes == true {
		t.Error("Contains should return false on Empty stack")
	}

	st.Push(1)

	yes = st.Contains(1)
	if yes == false {
		t.Error("Contains Got - false,expecting - true, after Push of 1")
	}

	st.Pop()
	yes = st.Contains(1)
	if yes == true {
		t.Error("Contains Got - true,expecting - false, after Pop of 1")
	}

}

func TestClear(t *testing.T) {
	st := NewStack()

	st.Push(1)

	st.Clear()
	cnt := st.Count()
	if cnt > 0 {
		t.Error("Clear Got - %d,expecting - 0, after clear of stack")
	}

}

func TestCount(t *testing.T) {
	st := NewStack()

	cnt := st.Count()
	if cnt != 0 {
		t.Errorf("Count() == %d, want %d", cnt, 0)
	}

	st.Push(1)
	cnt = st.Count()
	if cnt != 1 {
		t.Errorf("Count() == %d,want %d", cnt, 1)
	}

	st.Pop()
	cnt = st.Count()
	if cnt != 0 {
		t.Errorf("Count() == %d,want %d", cnt, 0)
	}

}
