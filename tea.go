package aqualuck

import "fmt"

// Tea is a concrete subclass that extends the template method
type Tea struct {
	CaffeineBeverage
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Adding lemon")
}
