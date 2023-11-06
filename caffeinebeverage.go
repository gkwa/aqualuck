package aqualuck

import "fmt"

// CaffeineBeverage is the concrete class that implements the template method
type CaffeineBeverage struct{}

func (b *CaffeineBeverage) BoilWater() {
	fmt.Println("Boiling water")
}

func (b *CaffeineBeverage) PourInCup() {
	fmt.Println("Pouring into cup")
}

func (b *CaffeineBeverage) Brew() {
	// This method should be overridden by concrete subclasses
}

func (b *CaffeineBeverage) AddCondiments() {
	// This method should be overridden by concrete subclasses
}

func (b *CaffeineBeverage) PrepareBeverage() {
	b.BoilWater()
	b.Brew() // Call the overridden Brew method
	b.PourInCup()
	b.AddCondiments() // Call the overridden AddCondiments method
}
