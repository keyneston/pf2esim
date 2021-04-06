package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/keyneston/pf2esim"
	pf "github.com/keyneston/pf2esim"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	d := pf.Dice{
		Type:   pf.DTPiercing,
		Number: 1,
		Size:   8,
	}

	deadlyD10 := pf.Dice{
		Type:   pf.DTPiercing,
		Number: 1,
		Size:   10,
	}

	att := pf2esim.Attack{
		Name:        "Bow",
		AttackBonus: 7,
		DamageDice:  append([]pf.Dice{}, d),
		CritDice:    append([]pf.Dice{}, deadlyD10),
	}

	for i := 0; i < 10; i++ {
		s := att.Calc()
		log.Printf("Dice: %q: %v to hit; %v dmg, %v on crit", att.Name, s.Attack, s.TotalDamage(), s.TotalCrit())
	}
}
