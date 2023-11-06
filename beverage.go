package aqualuck

// BeverageTemplate is the template method that defines the algorithm
type BeverageTemplate interface {
	BoilWater()
	PourInCup()

	Brew()
	AddCondiments()

	PrepareBeverage()
}
