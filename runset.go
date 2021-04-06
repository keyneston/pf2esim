package pf2esim

import (
	"fmt"
)

type RunSet struct {
	Name    string
	Attacks []Attack

	Results [][]AttackSummary
}

func (rs RunSet) Run(count int) {
	hits := make([]int, 0, count)
	dmgs := make([]int, 0, count)
	crits := make([]int, 0, count)

	for i := 0; i < count; i++ {
		for _, att := range rs.Attacks {
			s := att.Calc()

			hits = append(hits, s.Attack)
			dmgs = append(dmgs, s.TotalDamage())
			crits = append(crits, s.TotalCrit())
		}
	}

	fmt.Printf("# %s: Over %d Turns\n  Average Hit: %0.2f\n  Average Damage: %0.2f\n  Average Crit: %0.2f\n",
		rs.Name,
		count,
		avg(hits...),
		avg(dmgs...),
		avg(crits...),
	)
}

func avg(inputs ...int) float32 {
	s := 0
	for _, i := range inputs {
		s += i
	}

	return float32(s) / float32(len(inputs))
}
