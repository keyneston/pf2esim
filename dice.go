package pf2esim

import (
	"fmt"
	"math/rand"
)

type Dice struct {
	Number int
	Size   int
	Type   DamageType
	Bonus  int
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
