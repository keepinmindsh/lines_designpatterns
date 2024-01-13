package usecase

import "message/domain"

type publicChannel struct {
	message domain.Message
}

func newPublicChannel(message domain.Message) domain.Caller {
	return &publicChannel{
		message: message,
	}
}

func (p publicChannel) Send() {
	//TODO implement me
	panic("implement me")
}
