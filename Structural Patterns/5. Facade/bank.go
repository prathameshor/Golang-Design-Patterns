package main

type Bank struct {
}

func (b *Bank) Deposite(account *Account, amount float32) {
	account.balance += amount
}

func (b *Bank) Withdrawl(account *Account, amount float32) {
	account.balance -= amount
}

func (b *Bank) Transfer(source *Account, target *Account, amount float32) {
	source.balance -= amount
	target.balance += amount
}

func (b *Bank) StoreAmount(store Store, amount float32) {
	b.Deposite(store.(*Account), amount)
}
