package data

import (
	"clean_architecture/feature/logistic/domain"

	"gorm.io/gorm"
)

type LogisticModel struct {
	gorm.Model
	NamaVen  string
	Expedisi string
}

func FromDomain(dv domain.Logistic) LogisticModel {
	return LogisticModel{
		NamaVen:  dv.NamaVen,
		Expedisi: dv.Expedisi,
	}
}

func ToDomain(vm LogisticModel) domain.Logistic {
	return domain.Logistic{
		NamaVen:  vm.NamaVen,
		Expedisi: vm.Expedisi,
	}
}

func ToDomainArray(av []LogisticModel) []domain.Logistic {
	var res []domain.Logistic
	for _, val := range av {
		res = append(res, domain.Logistic{ID: val.ID, NamaVen: val.NamaVen, Expedisi: val.Expedisi})
	}
	return res
}
