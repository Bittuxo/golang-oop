package accounts

import (
	"golang-oop/clients"
)

type SaveAccount struct {
	Holder                     clients.Holder
	Agency, Account, Operation int
	save                       float64
}

func (a *SaveAccount) RemoveValue(value float64) string {
	canTake := value > 0 && value <= a.save
	if canTake {
		a.save -= value
		return "cash remover form account"
	} else {
		return "not enough cash"
	}
}

func (a *SaveAccount) AddValue(value float64) (string, float64) {
	canAdd := value > 0
	if canAdd {
		a.save += value
		return "cash add to account", a.save
	} else {
		return "wrong number", a.save
	}
}

func (a *SaveAccount) ShowMoney() float64 {
	return a.save
}
