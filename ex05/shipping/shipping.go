package shipping

import (
	"ex05/coubang"
)

type Sender interface {
	Send()
}

func Shipping(sender Sender) {
	c, ok := sender.(*coubang.Coubang)
	if ok {
		if c.Membership() {
			c.RocketSend()
			return
		}
	}
	sender.Send()
}
