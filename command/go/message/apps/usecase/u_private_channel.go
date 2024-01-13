package usecase

import "message/domain"

type privateChannel struct {
	message domain.Message
}

func newPrivateChannel(message domain.Message) domain.Caller {
	return &privateChannel{
		message: message,
	}
}

func (p privateChannel) Send() {
	//TODO implement me
	panic("implement me")
}
