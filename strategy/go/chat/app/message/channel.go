package message

import "chat/domain"

type channelMessage struct{}

func (c channelMessage) Create() {
	//TODO implement me
	panic("implement me")
}

func (c channelMessage) Notify() {
	//TODO implement me
	panic("implement me")
}

func (c channelMessage) Send() {
	//TODO implement me
	panic("implement me")
}

func NewChannelMessage() domain.Message {
	return channelMessage{}
}
