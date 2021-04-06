package pf2esim

import (
	"fmt"
)

type RunSet struct {
	Name    string
	Attacks []Attack
}

func (rs RunSet) Run(dc int, count int) {
	hitAttempts := 0
	hits := make([]int, 0, count)
	dmgs := make([]int, 0, count)
	crits := make([]int, 0, count)
	dmgPerTurn := make([]int, 0, count)
	critPerTurn := make([]int, 0, count)

	for i := 0; i < count; i++ {
		turnDmg := 0
		turnCrit := 0

		for _, att := range rs.Attacks {
			hitAttempts += 1

			s := att.Calc()

			hits = append(hits, s.Attack)
			if s.Attack >= dc {
				turnDmg += s.TotalDamage()
				dmgs = append(dmgs, s.TotalDamage())

			}

			if s.Attack >= dc+10 {
				turnCrit += s.TotalCrit()
				crits = append(crits, s.TotalCrit())
			}
		}

		dmgPerTurn = append(dmgPerTurn, turnDmg)
		critPerTurn = append(critPerTurn, turnCrit)
	}

	fmt.Printf("# %s: Over %d Turns\n  Hit Rate: %0.2f%%\n  Average Hit: %0.2f\n  Average Damage Per Hit: %0.2f\n  Average Crit Per Hit: %0.2f\n  Average Damage Per Turn: %0.2f\n  Average Crit per Turn: %0.2f\n",
		rs.Name,
		count,
		float32(len(dmgs))/float32(hitAttempts)*100,
		avg(hits...),
		avg(dmgs...),
		avg(crits...),
		avg(dmgPerTurn...),
		avg(critPerTurn...),
	)
}

func avg(inputs ...int) float32 {
	s := 0
	for _, i := range inputs {
		s += i
	}

	return float32(s) / float32(len(inputs))
}
