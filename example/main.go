package main

import (
	"math/rand"
	"time"

	pf "github.com/keyneston/pf2esim"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	att := pf.Attack{
		Name:        "Bow",
		AttackBonus: 7,
		DamageDice: []pf.Dice{
			pf.D8.SetType(pf.DTPiercing),
		},
		CritDice: []pf.Dice{
			pf.D10.SetType(pf.DTPiercing),
		},
	}

	bowRS := pf.RunSet{
		Name:    "2 x Bow Shots",
		Attacks: []pf.Attack{att, att.Second()},
	}

	alch := pf.Attack{
		Name:        "Alchemical Crossbow",
		AttackBonus: 7,
		DamageDice: []pf.Dice{
			pf.D8.SetType(pf.DTPiercing),
			pf.D6.SetType(pf.DTFire),
		},
	}

	alchRS := pf.RunSet{
		Name:    alch.Name,
		Attacks: []pf.Attack{alch},
	}

	bowRS.Run(15, 10000)
	alchRS.Run(15, 10000)

	//for i := 0; i < 10; i++ {
	//	s := att.Calc()
	//	log.Printf("Dice: %q: %v to hit; %v dmg, %v on crit", att.Name, s.Attack, s.TotalDamage(), s.TotalCrit())
	//}
}
