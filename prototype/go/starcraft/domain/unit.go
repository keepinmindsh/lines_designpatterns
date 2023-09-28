package domain

type Unit interface {
	Harvest()
	Attack()
	Building()
	GerMineralCapacity() int
}
