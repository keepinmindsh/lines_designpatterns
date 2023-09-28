package unit

import (
	"fmt"
	"starcraft/domain"
)

type Probe struct {
	MineralCapacity    int
	AccumulatedMineral int
}

func (p *Probe) Harvest() {
	fmt.Println("자원을 캡니다.")

	if p.MineralCapacity >= p.AccumulatedMineral {
		p.AccumulatedMineral += 1
	} else {
		fmt.Println("이미 캘수 있는 모든 미네랄을 다 캐었기 때문에 더이상 캘 수 없습니다.")
	}
}

func (p *Probe) Attack() {
	fmt.Println("공격을 합니다.")
}

func (p *Probe) Building() {
	fmt.Println("건물을 짓습니다.")
}

func (p *Probe) GerMineralCapacity() int {
	return p.MineralCapacity
}

func NewCloneUnit() domain.Unit {
	cloneObj := &Probe{
		MineralCapacity:    50,
		AccumulatedMineral: 0,
	}

	return cloneObj
}
