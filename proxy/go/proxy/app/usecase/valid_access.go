package usecase

import (
	"fmt"
	"proxy/domain"
)

type ValidAccess struct {
	UnAuthorizedAccess domain.UnAuthorizedAccess
}

func (v ValidAccess) Access() {
	v.UnAuthorizedAccess.NotValidAccess()

	fmt.Println("나는 인가 받았기에 언제든지 접속할 수 있는 권한이 있다.")
}

func NewValidAccess(access domain.UnAuthorizedAccess) domain.AuthorizedAccess {
	return ValidAccess{
		UnAuthorizedAccess: access,
	}
}
