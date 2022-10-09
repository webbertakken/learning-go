package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyer interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyer
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

type ToxicWaste struct {
	weight int
}

func (t ToxicWaste) String() string {
	return "Toxic waste"
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func (t *ToxicWaste) Convey() BeltSize {
	return Medium
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v conveyor and ship it via %v.\n", item, item.Convey(), item.Ship())
}

func main() {
	mail := SpamMail{40000}
	automate(&mail)

	waste := ToxicWaste{300}
	automate(&waste)
}
