package aqualuck

import "fmt"

// Coffee is another concrete subclass that extends the template method
type Coffee struct {
	CaffeineBeverage
}

func (c *Coffee) Brew() {
	fmt.Println("Dripping coffee through filter")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Adding sugar and milk")
}
