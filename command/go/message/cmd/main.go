package main

import (
	"message/apps/usecase"
	"message/domain"
)

func main() {
	messageCreate := usecase.NewMessageCreate()

	publicCaller := usecase.MustNewCaller(messageCreate, domain.PUBLIC_CHANNEL)
	publicCaller.Send()

	privateCaller := usecase.MustNewCaller(messageCreate, domain.PRIVATE_CHANNEL)
	privateCaller.Send()
}
