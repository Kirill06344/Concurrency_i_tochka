package main

import (
	"fmt"
	"sync"
)

type VkusnoITochka struct {
	burgers     int
	frenchFries int
}

func (vt *VkusnoITochka) GiveBurger() {
	vt.burgers--
}

func (vt *VkusnoITochka) GiveFrenchFries() {
	vt.frenchFries--
}

func (vt *VkusnoITochka) CookBurger() {
	vt.burgers++
}

func (vt *VkusnoITochka) CookFrenchFries() {
	vt.frenchFries++
}

func (vt *VkusnoITochka) ShowStocks() {
	fmt.Println("burgers: ", vt.burgers)
	fmt.Println("frenchFries: ", vt.frenchFries)
}

func VTSwampWithWork() {
	vt := VkusnoITochka{
		burgers:     10,
		frenchFries: 10,
	}

	fmt.Println("before work")
	vt.ShowStocks()

	var wg sync.WaitGroup
	wg.Add(40)
	orderBurger := func() {
		defer wg.Done()
		vt.GiveBurger()
	}

	orderFrenchFries := func() {
		defer wg.Done()
		vt.GiveFrenchFries()
	}

	cookBurger := func() {
		defer wg.Done()
		vt.CookBurger()
	}

	cookFrenchFries := func() {
		defer wg.Done()
		vt.CookFrenchFries()
	}

	for i := 0; i < 10; i++ {
		go orderBurger()
		go cookBurger()
		go orderFrenchFries()
		go cookFrenchFries()
	}

	wg.Wait()
	fmt.Println()
	fmt.Println("after this busy day...")
	vt.ShowStocks()
}
