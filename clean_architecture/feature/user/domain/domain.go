package domain

type User struct {
	ID   uint
	Nama string `json:"nama" form:"nama"`
	HP   string `json:"hp" form:"hp"`
}

type UserData interface {
	Insert(newUser User) (User, error)
	GetAll() ([]User, error)
	GetUser(Nama, HP string) (User, error)
}

type UserService interface {
	Register(newUser User) (User, error)
	ShowAllUser() ([]User, error)
	Login(Nama, HP string) (User, error)
}
