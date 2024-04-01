package main

type Store interface {
	SetBalance(float32)
	GetBalance() float32
}

type FinStoreage interface {
	StoreAmount(store Store, amount float32)
}

type FinStoreageFacade struct {
	storage FinStoreage
}

func NewFinstorageFacade(storage FinStoreage) *FinStoreageFacade {
	return &FinStoreageFacade{
		storage: storage,
	}
}

func (f *FinStoreageFacade) StoreAmount(store Store, amount float32) {
	f.storage.StoreAmount(store, amount)
}

