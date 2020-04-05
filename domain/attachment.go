package domain

import (
	"os"
	"time"
)

// Attachment domain describe file attached to another entity
type Attachment struct {
	ID        int
	Name      string
	File      os.File
	CreatedBy User
	CreatedAt time.Time
}
