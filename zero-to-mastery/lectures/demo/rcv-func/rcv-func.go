package main

import "log"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func vacateSpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkingLot{spaces: make([]Space, 3)}
	log.Println("initial:", lot)

	occupySpace(&lot, 2)
	lot.occupySpace(1)
	log.Println("after occupying some spaces:", lot)

	lot.vacateSpace(2)
	log.Println("after vacating space 2:", lot)
}
