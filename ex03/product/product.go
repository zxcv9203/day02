package product

import "errors"

type Product struct {
	price uint
	name  string
}

func NewProduct(price uint, name string) (*Product, error) {
	p := Product{}
	err := p.SetPrice(price)
	if err != nil {
		return nil, err
	}
	err = p.SetName(name)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (p *Product) SetPrice(price uint) error {
	if price == 0 {
		return errors.New("잘못된 가격이 입력되었습니다")
	}
	p.price = price
	return nil
}

func (p *Product) SetName(name string) error {
	if name == "" {
		return errors.New("잘못된 이름이 입력되었습니다")
	}
	p.name = name
	return nil
}

func (p *Product) Price() uint {
	return p.price
}

func (p *Product) Name() string {
	return p.name
}
