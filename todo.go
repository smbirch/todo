package todo

import "time"

type item struct {
	Task        string `json:"task"`
	Done        bool   `json:"done"`
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item
