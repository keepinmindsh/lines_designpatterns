package access

import (
	"fmt"
	"proxy/domain"
)

type Acceptor struct {
}

func (a Acceptor) Accept(access domain.AuthorizedAccess) {
	access.Access()

	fmt.Println("해당 접근이 허용되는 것을 확인하였습니다.")
}

func NewAcceptor() domain.Acceptor {
	return Acceptor{}
}
