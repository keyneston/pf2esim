package pf2esim

type Attack struct {
	Name string

	Agile       bool
	AttackBonus int
	DamageDice  []Dice
	CritDice    []Dice
	DamageBonus int
	BonusType   DamageType
}

func (a Attack) Calc() AttackSummary {
	summary := AttackSummary{
		Name:       a.Name,
		Attack:     D20.Roll(a.AttackBonus),
		Damage:     map[DamageType]int{},
		CritDamage: map[DamageType]int{},
	}

	for _, d := range a.DamageDice {
		dmg, t := d.RollTyped()

		summary.Damage[t] += dmg
		summary.CritDamage[t] += 2 * dmg
	}

	for _, d := range a.CritDice {
		dmg, t := d.RollTyped()

		summary.CritDamage[t] += dmg
	}

	return summary
}

func (a Attack) Second() Attack {
	if a.Agile {
		a.AttackBonus -= 4
	} else {
		a.AttackBonus -= 5
	}
	return a
}

func (a Attack) Third() Attack {
	if a.Agile {
		a.AttackBonus -= 8
	} else {
		a.AttackBonus -= 10
	}
	return a
}

type AttackSummary struct {
	Name       string
	Attack     int
	Damage     map[DamageType]int
	CritDamage map[DamageType]int
}

type DamageSet map[DamageType]int

func (ds DamageSet) Sum() int {
	sum := 0
	for _, i := range ds {
		sum += i
	}

	return sum
}

func (as AttackSummary) TotalDamage() int {
	return DamageSet(as.Damage).Sum()
}

func (as AttackSummary) TotalCrit() int {
	return DamageSet(as.CritDamage).Sum()
}
