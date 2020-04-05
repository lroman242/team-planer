package event

import "github.com/lroman242/team-planer/domain"

// Factory interface describe application events factory
// This instance should be used to simplify events creating
type Factory interface {
	NewUserRegisteredEvent(user *domain.User) Event
	NewUserPasswordUpdatedEvent(user *domain.User) Event
}
