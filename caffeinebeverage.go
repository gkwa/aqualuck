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

func (t *CaffeineBeverage) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *CaffeineBeverage) AddCondiments() {
	fmt.Println("Adding lemon")
}

func (b *CaffeineBeverage) PrepareBeverage() {
	b.BoilWater()
	b.Brew()
	b.PourInCup()
	b.AddCondiments()
}
