package domain

import "time"

// Task domain is used to describe piece of work in workspace
type Task struct {
	ID             int
	Title          string
	Description    string
	Attachments    []Attachment
	Comments       []Comment
	SubTasks       []Task
	Workspace      Workspace
	Repeatable     bool
	RepeatDuration time.Duration
	DueDate        time.Time
	AssignedTo     User
	CreatedBy      User
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
