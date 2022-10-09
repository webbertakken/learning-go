package main

import "testing"

//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

func newPlayer() Player {
	return Player{
		"Player",
		100,
		100,
		0,
		30,
	}
}

func TestHealth(t *testing.T) {
	player := newPlayer()

	if player.Health != 100 {
		t.Errorf("Expected player to start with 100 health, got %v", player.Health)
	}

	if player.MaxHealth != 100 {
		t.Errorf("Expected player to have 100 maxHealth, got %v", player.MaxHealth)
	}

	player.takeDamage(10)
	if player.Health != 90 {
		t.Fatalf("Expected player to take damage")
	}

	player.heal(10)
	if player.Health != 100 {
		t.Fatalf("Expected player to not heal above 100%%")
	}

	player.takeDamage(player.MaxHealth + 1)
	if player.Health != 0 {
		t.Fatalf("Expected player to not drop below 0 health, got %v", player.Health)
	}
}

func TestEnergy(t *testing.T) {
	player := newPlayer()

	if player.Energy != 0 {
		t.Errorf("Expected player to start with 0 energy, got %v", player.Energy)
	}

	if player.MaxEnergy != 30 {
		t.Errorf("Expected player to have 30 maxEnergy, got %v", player.MaxEnergy)
	}

	player.recharge(10)
	if player.Energy != 10 {
		t.Fatalf("Expected player to recharge energy")
	}

	player.useEnergy(5)
	if player.Energy != 5 {
		t.Fatalf("Expected player to use energy")
	}

	player.recharge(player.MaxEnergy + 1)
	if player.Energy != 30 {
		t.Fatalf("Expected player energy not to go about maxEnergy, got %v", player.Energy)
	}

	player.useEnergy(player.MaxEnergy + 1)
	if player.Energy != 0 {
		t.Fatalf("Expected player energy to not go into deficit, got %v", player.Energy)
	}
}
