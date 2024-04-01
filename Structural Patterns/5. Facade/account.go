package main

type Account struct {
	balance float32
}

func (a *Account) GetBalance() float32 {
	return a.balance
}

func (a *Account) SetBalance(bal float32) {
	a.balance = bal
}
