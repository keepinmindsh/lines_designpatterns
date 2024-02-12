package plane

import (
	"fmt"
	"right_airplan/domain"
	"strconv"
	"time"
)

type Plane struct {
	Name      string
	Landed    bool
	IsLanding bool
	Strip     *domain.Strip
}

func (p *Plane) GetName() string {
	return p.Name
}

func (p *Plane) LandingEnable() bool {
	return p.IsLanding
}

func (p *Plane) SetLandingEnable(isLanding bool) {
	p.IsLanding = isLanding
}

func (p *Plane) Landing() {
	p.IsLanding = true
	fmt.Println(p.Name + "이 착륙 중 입니다. ")
	var appendString string
	for i := 10; i > 0; i-- {
		time.Sleep(time.Second * 1)
		appendString += strconv.Itoa(i) + " 미터 상공 > "
	}
	fmt.Println(p.Name + " : " + appendString + "착륙 완료")
}

func (p *Plane) Stop() {
	(*p.Strip).Complete()
	p.IsLanding = false
	p.Landed = true
	fmt.Println(p.Name + "가 착륙했습니다.")
}

func (p *Plane) PlaneLanded() bool {
	return p.Landed
}

func NewPlane(planeName string, strip *domain.Strip) *Plane {
	return &Plane{Name: planeName, Strip: strip, Landed: false, IsLanding: false}
}
