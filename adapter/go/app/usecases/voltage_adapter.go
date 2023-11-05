package usecases

import (
	"adapter/domain"
	"fmt"
)

type VoltageAdapter struct {
	Adapter domain.V220ElectricAdapter
}

func NewVoltageAdapter(v220 domain.V220ElectricAdapter) domain.V110ElectricAdapter {
	return &VoltageAdapter{Adapter: v220}
}

func (v VoltageAdapter) UseElectric() {

	fmt.Println("V110 볼트의 전압을 V220의 기기를 사용할 수 있게 전압을 조정합니다. ")

	v.Adapter.UseElectric()
}
