package message

import "chat/domain"

type directMessage struct{}

func (d directMessage) Create() {
	//TODO implement me
	panic("implement me")
}

func (d directMessage) Notify() {
	//TODO implement me
	panic("implement me")
}

func (d directMessage) Send() {
	//TODO implement me
	panic("implement me")
}

func NewDirectMessage() domain.Message {
	return directMessage{}
}
