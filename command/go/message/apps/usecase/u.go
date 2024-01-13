package usecase

import "message/domain"

func MustNewCaller(message domain.Message, callType domain.CallType) domain.Caller {
	switch callType {
	case domain.PRIVATE_CHANNEL:
		return newPrivateChannel(message)
	case domain.PUBLIC_CHANNEL:
		return newPublicChannel(message)
	default:
		panic("error, check your callType")
	}
}

type message struct{}

func NewMessageCreate() domain.Message {
	return &message{}
}
