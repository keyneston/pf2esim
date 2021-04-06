package pf2esim

type Turn struct {
	Actions []Attack
}

type DamageReport struct {
	Damage     int
	CritDamage int
	Type       DamageType
}

type Report struct {
	ToHit  int
	damage int
	crit   int
}

// TotalDamage returns a sum of all the damage, and the relating crit damage.
func (r Report) TotalDamage() (int, int) {
	return 0, 0
}

func (t Turn) Calc() []Report {
	for _, act := range t.Actions {
		act.Calc()

	}

	return nil
}
