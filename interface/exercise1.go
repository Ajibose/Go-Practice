package inter

type Maintainer interface {
	MonthlyCost() float64
}

type Car struct {
	Insurance float64
	Gas       float64
}

type Truck struct {
	Insurance float64
	Diesel    float64
	TollFees  float64
}

func (c Car) MonthlyCost() float64 {
	return c.Insurance + c.Gas
}

func (t Truck) MonthlyCost() float64 {
	return t.Insurance + t.Diesel + t.TollFees
}

func TotalFleetCost(fleets []Maintainer) float64 {
	var totalCost float64
	for _, fleet := range fleets {
		totalCost += fleet.MonthlyCost()
	}

	return totalCost
}
