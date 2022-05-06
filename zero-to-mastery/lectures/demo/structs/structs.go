package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
		frog = Passenger{Name: "Frog", TicketNumber: 4}
	)
	fmt.Println(bill, ella, frog)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 5
	fmt.Println(heidi)

	bill.Boarded = true
	frog.Boarded = true

	if bill.Boarded {
		fmt.Println(bill.Name, "has boarded")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
