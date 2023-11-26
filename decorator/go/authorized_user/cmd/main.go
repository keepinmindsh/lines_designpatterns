package main

import (
	"authorized_user/app/usecase"
	"fmt"
)

func main() {
	uCase := usecase.NewUserAuthUCase(usecase.NewUser())

	user := uCase.GetAuthorizedUser()

	fmt.Println(user)
}
