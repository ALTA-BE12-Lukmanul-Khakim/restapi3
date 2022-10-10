package delivery

import "clean_architecture/feature/user/domain"

type RegisterFormat struct {
	Nama string `json:"nama" form:"nama"`
	HP   string `json:"hp" form:"hp"`
}

type LoginFormat struct {
	Nama string `json:"nama" form:"nama"`
	HP   string `json:"hp" form:"hp"`
}

func ToDomain(i interface{}) domain.User {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.User{Nama: cnv.Nama, HP: cnv.HP}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.User{Nama: cnv.Nama, HP: cnv.HP}
	}
	return domain.User{}
}
