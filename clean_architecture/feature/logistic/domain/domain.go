package domain

type Logistic struct {
	ID       uint
	NamaVen  string
	Expedisi string
}

type LogisticData interface {
	Insert(newLogistic Logistic) (Logistic, error)
	GetAll() ([]Logistic, error)
	GetVen(expedisi string) ([]Logistic, error)
}

type LogisticService interface {
	AddLogistic(newLogistic Logistic) (Logistic, error)
	ShowAllLogistic() ([]Logistic, error)
	ShowExpLogistic(expedisi string) ([]Logistic, error)
}
