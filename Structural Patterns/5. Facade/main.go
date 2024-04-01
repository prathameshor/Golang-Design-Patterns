package main

func main() {
	var finOps = NewFinstorageFacade(&Bank{})
	finOps.StoreAmount(&Account{}, 125.50)
}
