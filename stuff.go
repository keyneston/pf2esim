package pf2esim

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
