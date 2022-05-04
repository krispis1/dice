package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NumberOfDice     = 2
	RollCount        = 1
	NumberOfSides    = 6 // Uses NumberOfSides by default.
	MaxNumberOfSides = 0 // If MaxNumberOfSides >= MinNumberOfSides, each die will have a randomly generated number of sides [MinNumberOfSides; MaxNumberOfSides]

	MinNumberOfSides = 2 // should remain unchanged. Die should have at least two sides!
)

type dice struct {
	numberOfSides int
	rolledValue   int
}

type diceList []dice

func initDice(numberOfDice int) diceList {
	var diceList diceList
	var numberOfSides int

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numberOfDice; i++ {
		if MaxNumberOfSides >= MinNumberOfSides {
			numberOfSides = rand.Intn(MaxNumberOfSides-1) + 2 // min number of sides = 2
		} else {
			numberOfSides = NumberOfSides
		}

		diceList = append(diceList, dice{numberOfSides: numberOfSides})
	}

	return diceList
}

func (d *diceList) roll(numberOfRolls int) {
	for i := 1; i <= numberOfRolls; i++ {
		fmt.Printf("Roll #%d\n", i)

		for j := 0; j < len(*d); j++ {
			(*d)[j].rolledValue = rand.Intn((*d)[j].numberOfSides) + 1 // min rolled value = 1
		}

		d.print()
	}
}

func (d diceList) print() {
	var totalRoll int

	for _, dice := range d {
		totalRoll += dice.rolledValue
	}

	fmt.Println("Total roll:", totalRoll)

	if len(d) == 2 && totalRoll == 2 {
		fmt.Println("Snake eyes")
	}

	if totalRoll == 7 {
		fmt.Println("Lucky 7")
	}
}
