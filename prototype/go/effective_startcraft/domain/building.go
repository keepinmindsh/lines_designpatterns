package domain

type Building interface {
	CreateUnit(action UnitAction) Unit
}
