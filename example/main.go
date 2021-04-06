package main

import (
	"math/rand"
	"time"

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

	att := pf.Attack{
		Name:        "Bow",
		AttackBonus: 7,
		DamageDice:  append([]pf.Dice{}, d),
		CritDice:    append([]pf.Dice{}, deadlyD10),
	}

	bowRS := pf.RunSet{
		Name:    "2 x Bow Shots",
		Attacks: []pf.Attack{att, att.Second()},
	}

	alch := pf.Attack{
		Name:        "Alchemical Crossbow",
		AttackBonus: 7,
		DamageDice: []pf.Dice{d, pf.Dice{
			Type:   pf.DTFire,
			Number: 1,
			Size:   6,
		},
		},
	}

	alchRS := pf.RunSet{
		Name:    alch.Name,
		Attacks: []pf.Attack{alch},
	}

	bowRS.Run(10000)
	alchRS.Run(10000)

	//for i := 0; i < 10; i++ {
	//	s := att.Calc()
	//	log.Printf("Dice: %q: %v to hit; %v dmg, %v on crit", att.Name, s.Attack, s.TotalDamage(), s.TotalCrit())
	//}
}
