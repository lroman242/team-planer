package interactor

import (
	"errors"
	"github.com/lroman242/team-planer/domain"
	"github.com/lroman242/team-planer/usecase/event"
	"github.com/lroman242/team-planer/usecase/repository"
	"github.com/lroman242/team-planer/usecase/tool"
	"strings"
)

// userInteractor implement core functions related to the user
type userInteractor struct {
	repository   repository.UserRepository
	validator    tool.Validator
	hasher       tool.Hasher
	logger       tool.Logger
	eventChannel chan<- event.Event
	eventFactory event.Factory
}

// NewUserInteractor function helps to create UserInteractor instance
func NewUserInteractor(r repository.UserRepository, v tool.Validator, h tool.Hasher, l tool.Logger, eChannel chan <- event.Event, eFactory event.Factory) UserInteractor {
	return &userInteractor{
		repository: r,
		validator: v,
		hasher: h,
		logger: l,
		eventChannel: eChannel,
		eventFactory: eFactory,
	}
}

// UserInteractor contains all business logic related to the user instance
type UserInteractor interface {
	// RegisterUserWithEmail register new user with email & password credentials
	RegisterUserWithEmail(name, email, passowrd string) (*domain.User, error)

	// ChangePassword change user password
	ChangePassword(email, oldPassword, newPassword string) error
}

// RegisterUserWithEmail register new user with email & password credentials
// Validation rules:
// email - valid email address, unique
// name - min:2, max:55
// password - min:6, max:255
func (i *userInteractor) RegisterUserWithEmail(name, email, password string) (*domain.User, error) {
	err := i.validator.IsEmail(email)
	if err != nil {
		return nil, errors.New("invalid email provided")
	}

	if i.repository.IsEmailExists(email) {
		return nil, errors.New("email is already used")
	}

	if len(name) < 2 {
		return nil, errors.New("name is too short. expected at least 3 chars")
	}

	if len(name) > 55 {
		return nil, errors.New("name is too long. expected max 55 chars")
	}

	if len(password) > 255 {
		return nil, errors.New("password is too long. expected max 255 chars")
	}

	if len(password) < 6 {
		return nil, errors.New("password is too short. expected at least 6 chars")
	}

	hash, err := i.hasher.Hash(password)
	if err != nil {
		i.logger.Printf("an error occurred during hashing password. error: %s", err)
		return nil, errors.New("unable to hash password. an error occurred")
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: hash,
	}

	err = i.repository.Create(user)
	if err != nil {
		i.logger.Printf("an error occurred during storing user into database. error: %s", err)
		return nil, errors.New("unable to create user")
	}

	// trigger UserRegisteredEvent
	i.eventChannel <- i.eventFactory.NewUserRegisteredEvent(user)

	return user, nil
}

// ChangePassword change user password
func (i *userInteractor) ChangePassword(email, oldPassword, newPassword string) error {
	err := i.validator.IsEmail(email)
	if err != nil {
		return errors.New("invalid email provided")
	}

	user, err := i.repository.FindByEmail(email)
	if err != nil {
		return errors.New("user with such email not found")
	}

	if !i.hasher.Compare(oldPassword, user.Password) {
		return errors.New("wrong old password provided")
	}

	if len(newPassword) > 255 {
		return errors.New("password is too long. expected max 255 chars")
	}

	if len(newPassword) < 6 {
		return errors.New("password is too short. expected at least 6 chars")
	}

	if strings.Compare(oldPassword, newPassword) == 0 {
		return errors.New("old password is same as new")
	}

	user.Password, err = i.hasher.Hash(newPassword)
	if err != nil {
		return errors.New("unable to hash password. an error occurred")
	}

	err = i.repository.Update(user)
	if err != nil {
		return errors.New("unable to update password. an error occurerd")
	}

	// trigger userPasswordUpdatedEvent
	i.eventChannel <- i.eventFactory.NewUserPasswordUpdatedEvent(user)

	return nil
}
