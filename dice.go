package pf2esim

import (
	"fmt"
	"math/rand"
)

var (
	D20 = Dice{
		Number: 1,
		Size:   20,
		Type:   DTUntyped,
	}
	D12 = Dice{
		Number: 1,
		Size:   12,
		Type:   DTUntyped,
	}
	D10 = Dice{
		Number: 1,
		Size:   10,
		Type:   DTUntyped,
	}
	D8 = Dice{
		Number: 1,
		Size:   8,
		Type:   DTUntyped,
	}
	D6 = Dice{
		Number: 1,
		Size:   6,
		Type:   DTUntyped,
	}
	D4 = Dice{
		Number: 1,
		Size:   4,
		Type:   DTUntyped,
	}
)

type Dice struct {
	Number int
	Size   int
	Type   DamageType
	Bonus  int
}

func (d Dice) SetType(t DamageType) Dice {
	d.Type = t
	return d
}

func (d Dice) IncreaseSize() Dice {
	switch d.Size {
	case 10:
		d.Size = 12
	case 8:
		d.Size = 10
	case 6:
		d.Size = 8
	case 4:
		d.Size = 6
	}
	return d
}

func (d Dice) DecreaseSize() Dice {
	switch d.Size {
	case 12:
		d.Size = 10
	case 10:
		d.Size = 8
	case 8:
		d.Size = 6
	case 6:
		d.Size = 4
	}
	return d
}

func (d Dice) RollTyped() (int, DamageType) {
	return d.Roll(), d.Type
}

func (d Dice) Roll(bonuses ...int) int {
	sum := 0
	sumBonus := d.Bonus

	for _, bonus := range bonuses {
		sumBonus += bonus
	}

	for i := 0; i < d.Number; i++ {
		// Add 1 to move it from (for a d8) 0-7 to 1-8
		sum += ((rand.Int() % d.Size) + 1) + sumBonus
	}

	return sum
}

func (d Dice) String() string {
	return fmt.Sprintf("%dd%d+%d %s", d.Number, d.Size, d.Bonus, d.Type)
}

func ParseDice(in string) (Dice, error) {
	return Dice{}, fmt.Errorf("Not yet implemented")
}

func MustParseDice(in string) Dice {
	d, err := ParseDice(in)
	if err != nil {
		panic(err)
	}

	return d
}
