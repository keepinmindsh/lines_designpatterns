package usecases

import (
	"adapter/domain"
	"fmt"
)

type V110 struct {
}

func (v V110) UseElectric() {
	fmt.Println("110 볼트의 전기 제품을 사용할 수 있습니다.")
}

func NewV110() domain.V110ElectricAdapter {
	return &V110{}
}
