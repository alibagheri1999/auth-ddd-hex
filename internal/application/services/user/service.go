package user

type UserService interface {
	CreateUser(name, email, password string) error
}
