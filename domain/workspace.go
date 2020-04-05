package domain

import "time"

// Workspace domain is used to combine tasks related to one subject/project
type Workspace struct {
	ID        int
	Name      string
	Owner     User
	Users     []User
	CreatedAt time.Time
	UpdatedAt time.Time
}
