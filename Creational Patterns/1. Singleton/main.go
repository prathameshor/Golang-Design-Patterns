package main

import (
	"sync"

	"github.com/prathameshor/design-patterns/singleton/single"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			single.GetInstance()
		}()
	}
	wg.Wait()
}
