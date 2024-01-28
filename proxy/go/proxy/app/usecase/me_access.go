package usecase

import (
	"fmt"
	"proxy/domain"
)

type Me struct {
}

func (m Me) NotValidAccess() {
	fmt.Println("나는 인가 받지 않았지만 접속을 요청한다.")
}

func NewMeAccess() domain.UnAuthorizedAccess {
	return Me{}
}
