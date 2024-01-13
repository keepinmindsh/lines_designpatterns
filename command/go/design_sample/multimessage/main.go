package main

func main() {

}

type Caller struct {
	message Message
}

func NewCaller(msg Message) *Caller {
	return &Caller{
		message: msg,
	}
}

func (c Caller) ChannelCall() {
	// 메세지를 전송한다는 명령은 동일하다.
	c.message.MessageCreate()
}

func (c Caller) DirectMessageCall() {
	// 메세지를 전송한다는 명령은 동일하다.
	c.message.MessageCreate()
}

type Message interface {
	MessageCreate()
}

type messageCreate struct {
}

func NewMessage() Message {
	return &messageCreate{}
}

// 명령을 수행하는 MessageCreate는 호출 이후의 상태 관리를 자체적으로 수행해야 한다.
func (m messageCreate) MessageCreate() {

}
