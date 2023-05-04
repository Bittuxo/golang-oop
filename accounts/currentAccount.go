package accounts

import (
	"golang-oop/clients"
)

type CurrentAccount struct {
	Holder          clients.Holder
	Agency, Account int
	balance         float64
}

func (a *CurrentAccount) RemoveValue(value float64) string {
	canTake := value > 0 && value <= a.balance
	if canTake {
		a.balance -= value
		return "cash remover form account"
	} else {
		return "not enough cash"
	}
}

func (a *CurrentAccount) AddValue(value float64) (string, float64) {
	canAdd := value > 0
	if canAdd {
		a.balance += value
		return "cash add to account", a.balance
	} else {
		return "wrong number", a.balance
	}
}

func (a *CurrentAccount) SendMoney(sendValue float64, accountDestination *CurrentAccount) bool {
	if sendValue < a.balance && sendValue > 0 {
		a.balance -= sendValue
		accountDestination.AddValue(sendValue)
		return true
	} else {
		return false
	}
}

func (a *CurrentAccount) ShowMoney() float64 {
	return a.balance
}
