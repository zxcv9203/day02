package coubang

import (
	"errors"
	"fmt"

	"ex05/product"
)

type Coubang struct {
	product.Product
	to         string
	from       string
	membership bool
}

func (c *Coubang) SetTo(name string) error {
	if name == "" {
		return errors.New("잘못된 이름이 들어왔습니다")
	}
	c.to = name
	return nil
}

func (c *Coubang) SetFrom(name string) error {
	if name == "" {
		return errors.New("잘못된 이름이 들어왔습니다")
	}
	c.from = name
	return nil
}

func (c *Coubang) SetMembership(membership bool) {
	c.membership = membership
}

func (c *Coubang) SetProduct(p product.Product) error {
	err := c.SetPrice(p.Price())
	if err != nil {
		return err
	}
	err = c.SetName(p.Name())
	if err != nil {
		return err
	}
	return nil
}

func (c *Coubang) To() string {
	return c.to
}

func (c *Coubang) From() string {
	return c.from
}

func (c *Coubang) Membership() bool {
	return c.membership
}

func (c *Coubang) Send() {
	ok := c.Info()
	if !ok {
		return
	}
	to := c.To()
	from := c.From()
	if to == "" {
		fmt.Println("택배 발송 전 발송자를 먼저 입력해주세요")
		return
	}
	if from == "" {
		fmt.Println("택배 발송 전 수신자를 먼저 입력해주세요")
		return
	}
	fmt.Printf("[COUBANG] %s 님이 %s 님에게 일반 배송을 시작했습니다.\n", to, from)
}

func (c *Coubang) Info() bool {
	name := c.Name()
	if name == "" {
		fmt.Println("상품명을 먼저 설정해주세요")
		return false
	}
	price := c.Price()
	if price == 0 {
		fmt.Println("상품의 가격을 먼저 설정해주세요")
		return false
	}
	fmt.Println("상품명 :", name)
	fmt.Printf("상품가격 : %d원\n", price)
	return true
}

func (c *Coubang) RocketSend() {
	ok := c.Info()
	if !ok {
		return
	}
	to := c.To()
	from := c.From()
	if to == "" {
		fmt.Println("택배 발송전 발송자를 먼저 입력해주세요")
		return
	}
	if from == "" {
		fmt.Println("택배 발송전 수신자를 먼저 입력해주세요")
		return
	}
	fmt.Printf("[COUBANG] %s 님이 %s 님에게 로켓 배송을 시작했습니다.\n", to, from)
}

func NewCoubang(p product.Product, to, from string, membership bool) (*Coubang, error) {
	coubang := Coubang{}
	err := coubang.SetProduct(p)
	if err != nil {
		return nil, err
	}
	err = coubang.SetTo(to)
	if err != nil {
		return nil, err
	}
	err = coubang.SetFrom(from)
	if err != nil {
		return nil, err
	}
	coubang.SetMembership(membership)
	return &coubang, nil
}
