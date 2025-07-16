package main

// Implement a ride fare calculator that can switch between pricing strategies: regular, night, and surge.
// Each strategy uses a different rate per kilometer and time of day.
// The code should support adding new strategies easily without touching core logic.

import (
	"fmt"
)

type FareStrategy interface {
	CalculateFare(distance float64) float64
}

// Night Fare
type NightFare struct {
	Multiplier int
}

func NewNightFare(multiplier int) *NightFare {
	return &NightFare{Multiplier: multiplier}
}

func (n *NightFare) CalculateFare(distance float64) float64 {
	return distance * float64(10*n.Multiplier)
}

// Surge Fare Strategy
type SurgeFactors struct {
	CityFactor   float64
	TimeFactor   float64
	DemandFactor float64
}

type SurgeFare struct {
	Factors SurgeFactors
}

func NewSurgeFare(factors SurgeFactors) *SurgeFare {
	return &SurgeFare{Factors: factors}
}

func (s *SurgeFare) CalculateFare(distance float64) float64 {
	effectiveRate := 10.0 * s.Factors.CityFactor * s.Factors.TimeFactor * s.Factors.DemandFactor
	return distance * effectiveRate
}

// Regular Fare Strategy

type RegularFare struct{}

func NewRegularFare() *RegularFare {
	return &RegularFare{}
}

func (r *RegularFare) CalculateFare(distance float64) float64 {
	return distance * 10.0
}

// Calculator

type FareCalculator struct {
	strategy FareStrategy
}

func NewFareCalculator(strategy FareStrategy) *FareCalculator {
	return &FareCalculator{strategy: strategy}
}

func (fc *FareCalculator) SetStrategy(strategy FareStrategy) {
	fc.strategy = strategy
}

func (fc *FareCalculator) Calculate(distance float64) float64 {
	return fc.strategy.CalculateFare(distance)
}

func main() {
	distance := 10.0

	// Regular Fare
	calculator := NewFareCalculator(NewRegularFare())
	fmt.Printf("Regular Fare: $%.2f\n", calculator.Calculate(distance))

	// Night Fare with multiplier
	calculator.SetStrategy(NewNightFare(2))
	fmt.Printf("Night Fare: $%.2f\n", calculator.Calculate(distance))

	// Surge Fare with factors
	surge := SurgeFactors{CityFactor: 1.2, TimeFactor: 1.5, DemandFactor: 1.3}
	calculator.SetStrategy(NewSurgeFare(surge))
	fmt.Printf("Surge Fare: $%.2f\n", calculator.Calculate(distance))
}
