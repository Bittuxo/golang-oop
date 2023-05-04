package main

import (
	"fmt"
	a "golang-oop/accounts"
	//c "golang-oop/clients"
)

func PayBills(account verifyAccount, valueBills float64) {
	account.RemoveValue(valueBills)
}

type verifyAccount interface {
	RemoveValue(value float64) string
}

func main() {
	DenisAccount := a.SaveAccount{}
	DenisAccount.AddValue(5000)
	PayBills(&DenisAccount, 500)

	fmt.Println(DenisAccount.ShowMoney())

	LuisaAccount := a.CurrentAccount{}
	LuisaAccount.AddValue(250)
	PayBills(&LuisaAccount, 150)

	fmt.Println(LuisaAccount.ShowMoney())
}
