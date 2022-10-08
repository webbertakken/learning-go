package main

import "log"

//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
type Player struct {
	name      string
	Health    int
	MaxHealth int
	Energy    int
	MaxEnergy int
}

//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

func (p *Player) takeDamage(amount int) {
	p.Health -= amount

	if p.Health < 0 {
		p.Health = 0
	}

	log.Printf("OUch! (health: %d)\n", p.Health)
}

func (p *Player) heal(amount int) {
	p.Health += amount

	if p.Health > p.MaxHealth {
		p.Health = p.MaxHealth
	}

	log.Printf("Aaahh! (health: %d)\n", p.Health)
}

func (p *Player) useEnergy(amount int) {
	p.Energy -= amount

	if p.Energy < 0 {
		p.Energy = 0
	}

	log.Printf("Slurpf. (energy: %d)\n", p.Energy)
}

func (p *Player) recharge(amount int) {
	p.Energy += amount

	if p.Energy > p.MaxEnergy {
		p.Energy = p.MaxEnergy
	}

	log.Printf("Floomph! (energy: %d)\n", p.Energy)
}

func main() {
	noice := Player{"Noice Doode", 100, 100, 0, 30}

	noice.heal(5)
	noice.recharge(9)
	noice.useEnergy(20)
	noice.recharge(40)
	noice.takeDamage(99)
	noice.takeDamage(5)
	noice.heal(999)
}
