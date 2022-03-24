package todo_test

import (
	"testing"
	"todo"
)

// TestAdd tests the Add method of the List type
func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead", taskName, l[0].Task)
	}
}

// TestComplete tests the Complete method of the List type
