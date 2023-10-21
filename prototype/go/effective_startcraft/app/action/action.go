package action

import (
	"effective_startcraft/domain"
	"fmt"
)

type Action struct {
	MineralCapacity int
}

func (p *Action) Harvest(accumulatedMineral int) int {
	fmt.Println("자원을 캡니다.")

	if p.MineralCapacity >= accumulatedMineral {
		accumulatedMineral += 1
	} else {
		fmt.Println("이미 캘수 있는 모든 미네랄을 다 캐었기 때문에 더이상 캘 수 없습니다.")
	}

	return accumulatedMineral
}

func (p *Action) Attack() {
	fmt.Println("공격을 합니다.")
}

func (p *Action) Building() {
	fmt.Println("건물을 짓습니다.")
}

func NewAction() domain.UnitAction {
	return &Action{
		MineralCapacity: 50,
	}
}
