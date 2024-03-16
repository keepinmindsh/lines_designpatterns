package chat

import (
	"chat/domain"
	"sync"
)

type chat struct {
	Context domain.Message
}

func (c *chat) MessageCreate() {
	c.Context.Create()

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		c.Context.Notify()
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		c.Context.Notify()
		wg.Done()
	}()

	wg.Wait()
}

func (c *chat) RunContext(message domain.Message) {
	c.Context = message
}

func NewChat() domain.Chat {
	return &chat{}
}
