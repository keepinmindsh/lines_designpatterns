package usecases

import (
	"adapter/domain"
	"fmt"
)

type V220 struct {
}

func (v V220) UseElectric() {
	fmt.Println("220 볼트의 전기 제품을 사용할 수 있습니다.")
}

func NewV220() domain.V220ElectricAdapter {
	return &V220{}
}
