package main

import (
	"fmt"
	"sync"
)

type Rostics struct {
	wings   int32
	nuggets int32
	mu      sync.RWMutex
}

func (rt *Rostics) GiveNuggets() {
	rt.mu.Lock()
	defer rt.mu.Unlock()
	rt.nuggets--
}

func (rt *Rostics) GiveWings() {
	rt.mu.Lock()
	defer rt.mu.Unlock()
	rt.wings--
}

func (rt *Rostics) CookWings() {
	rt.mu.Lock()
	defer rt.mu.Unlock()
	rt.wings++
}

func (rt *Rostics) CookNuggets() {
	rt.mu.Lock()
	defer rt.mu.Unlock()
	rt.nuggets++
}

func (rt *Rostics) ShowStocks() {
	rt.mu.RLock()
	defer rt.mu.RUnlock()
	fmt.Println("nuggets: ", rt.nuggets)
	fmt.Println("wings: ", rt.wings)
}

func RosticsSwampWithWork() {
	rt := Rostics{
		wings:   10,
		nuggets: 10,
	}

	fmt.Println("before work")
	rt.ShowStocks()

	var wg sync.WaitGroup
	wg.Add(40)
	orderWings := func() {
		defer wg.Done()
		rt.GiveWings()
	}

	orderNuggets := func() {
		defer wg.Done()
		rt.GiveNuggets()
	}

	cookWings := func() {
		defer wg.Done()
		rt.CookWings()
	}

	cookNuggets := func() {
		defer wg.Done()
		rt.CookNuggets()
	}

	for i := 0; i < 10; i++ {
		go orderWings()
		go cookNuggets()
		go cookWings()
		go orderNuggets()
	}

	wg.Wait()
	fmt.Println()
	fmt.Println("after this busy day...")
	rt.ShowStocks()
}
