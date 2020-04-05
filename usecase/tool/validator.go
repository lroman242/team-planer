package tool

type Validator interface {
	IsEmail(email string) error

}
