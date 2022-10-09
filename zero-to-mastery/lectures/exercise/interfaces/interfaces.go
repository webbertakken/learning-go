//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:

//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts

type Lift int

const (
	SmallLift Lift = iota
	StandardLift
	LargeLift
)

type LiftDirector interface {
	PickLift() Lift
}

type Vehicle struct {
	make  string
	model string
}

type Motorcycle struct {
	Vehicle
}

type Car struct {
	Vehicle
}

type Truck struct {
	Vehicle
}

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v, %v", m.make, m.model)
}

func (m Motorcycle) PickLift() Lift {
	return SmallLift
}

func (c Car) String() string {
	return fmt.Sprintf("Car: %v, %v", c.make, c.model)
}

func (c Car) PickLift() Lift {
	return StandardLift
}

func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v, %v", t.make, t.model)
}

func (t Truck) PickLift() Lift {
	return LargeLift
}

//* Write a single function to handle all of the vehicles
//  that the shop works on.
func MoveToLift(d LiftDirector) {
	switch d.PickLift() {
	case SmallLift:
		fmt.Printf("send %v to small lift\n", d)
	case StandardLift:
		fmt.Printf("send %v to standard lift\n", d)
	case LargeLift:
		fmt.Printf("send %v to large lift\n", d)
	}
}

type Lifts struct {
	motorcycleLift Motorcycle
	carLift        Car
	truckLift      Truck
}

func main() {
	incomingVehicles := []LiftDirector{Car{Vehicle{"Mazda", "323"}}, Motorcycle{Vehicle{"Kawasaki", "Ninja 400"}}, Truck{Vehicle{"IVECO", "Daily 4x4"}}}

	for _, vehicle := range incomingVehicles {
		MoveToLift(vehicle)
	}
}
