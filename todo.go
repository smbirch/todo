package todo

import (
	"fmt"
	"time"
)

type item struct {
	Task        string `json:"task"`
	Done        bool   `json:"done"`
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Now(),
	}

	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("task %d does not exist", i)
	}

	// Adjusting for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("task %d does not exist", i)
	}

	// Adjusting for 0 based index
	*l = append(ls[:i-1], ls[i:]...)

	return nil

}
