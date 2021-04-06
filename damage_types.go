package pf2esim

type DamageType string

const (
	DTUntyped     DamageType = "untyped"
	DTPiercing    DamageType = "piercing"
	DTSlashing               = "slashing"
	DTBludgeoning            = "bludgeoning"
	DTPrecision              = "precision"
)
