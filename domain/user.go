package domain

// User domain describe application user
type User struct {
	ID       int
	Name     string
	Email    string
	Password []byte
}
