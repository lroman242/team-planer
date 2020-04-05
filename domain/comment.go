package domain

import "time"

// Comment is domain used to describe messages attached to another entities
type Comment struct {
	ID        int
	Text      string
	Author    User
	Task      Task
	CreatedAt time.Time
	UpdatedAt time.Time
}
