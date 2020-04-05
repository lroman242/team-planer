package repository

import "github.com/lroman242/team-planer/domain"

type UserRepository interface {
	Create(user *domain.User) error
	IsEmailExists(email string) bool
	FindByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
}
