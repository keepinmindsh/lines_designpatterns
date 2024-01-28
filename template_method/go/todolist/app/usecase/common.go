package usecase

import (
	"fmt"
	"todolist/domain"
)

type Common struct {
}

func (c Common) MakeTitle() {
	fmt.Println("타이틀을 작성합니다.")
}

func (c Common) MakePeriod() {
	fmt.Println("기간을 입력합니다.")
}

func (c Common) MakeConfirm() {
	fmt.Println("확인 및 저장힙니다.")
}

func NewCommonToDo() domain.CommonToDo {
	return &Common{}
}
