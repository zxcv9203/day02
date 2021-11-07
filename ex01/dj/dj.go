package dj

import (
	"errors"
	"ex01/product"
	"fmt"
)

type Dj struct {
	product.Product
	to   string
	from string
}

func (d *Dj) SetTo(name string) error {
	if name == "" {
		return errors.New("잘못된 이름이 들어왔습니다")
	}
	d.to = name
	return nil
}

func (d *Dj) SetFrom(name string) error {
	if name == "" {
		return errors.New("잘못된 이름이 들어왔습니다")
	}
	d.from = name
	return nil
}

func (d *Dj) SetProduct(p product.Product) error {
	err := d.SetPrice(p.Price())
	if err != nil {
		return err
	}
	err = d.SetName(p.Name())
	if err != nil {
		return err
	}
	return nil
}

func (d *Dj) To() string {
	return d.to
}

func (d *Dj) From() string {
	return d.from
}

func NewDj(p product.Product, to, from string) (*Dj, error) {
	dj := Dj{}
	err := dj.SetProduct(p)
	if err != nil {
		return nil, err
	}
	err = dj.SetTo(to)
	if err != nil {
		return nil, err
	}
	err = dj.SetFrom(from)
	if err != nil {
		return nil, err
	}
	return &dj, nil
}

func (d *Dj) Info() bool {
	name := d.Name()
	if name == "" {
		fmt.Println("상품명을 먼저 설정해주세요")
		return false
	}
	price := d.Price()
	if price == 0 {
		fmt.Println("상품의 가격을 먼저 설정해주세요")
		return false
	}
	fmt.Println("상품명 :", name)
	fmt.Printf("상품가격 : %d원\n", price)
	return true
}

func (d *Dj) Send() {
	ok := d.Info()
	if !ok {
		return
	}
	to := d.To()
	from := d.From()
	if to == "" {
		fmt.Println("택배 발송 전 발송자를 먼저 입력해주세요")
		return
	}
	if from == "" {
		fmt.Println("택배 발송 전 수신자를 먼저 입력해주세요")
		return
	}
	fmt.Printf("[DJ] %s 님이 %s 님에게 일반 배송을 시작했습니다.\n", to, from)
}
