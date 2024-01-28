package main

import (
	"proxy/app/access"
	"proxy/app/usecase"
)

func main() {

	authorizedAccess := usecase.NewValidAccess(usecase.NewMeAccess())

	acceptor := access.NewAcceptor()

	acceptor.Accept(authorizedAccess)

	// 아래의 코드는 인가 받지 않음 함수이기 때문에 접근할 수 없다.
	//acceptor.Accept(usecase.NewMeAccess())
}
