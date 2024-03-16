package main

import (
	"chat/app/chat"
	"chat/app/message"
	"chat/domain"
)

func main() {
	runMessage(message.NewChannelMessage())

	runMessage(message.NewDirectMessage())
}

func runMessage(message domain.Message) {
	chatting := chat.NewChat()

	chatting.RunContext(message)

	chatting.MessageCreate()
}
