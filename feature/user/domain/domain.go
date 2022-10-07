package domain

type User struct {
	ID   uint
	Nama string
	HP   string
}

type UserData interface {
	Insert(newUser User) (User, error)
	GetAll() ([]User, error)
}

type UserService interface {
	AddUser(newUser User) (User, error)
	ShowAllUser() ([]User, error)
}
