package main

import (
	"ex00/product"
	"fmt"
)

func errorName(err error, name string) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(name)
	}
}

func errorPrice(err error, price uint) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(price)
	}
}

func main() {
	p := product.Product{}

	errorName(p.SetName(""), p.Name())
	errorName(p.SetName("candy"), p.Name())
	errorPrice(p.SetPrice(0), p.Price())
	errorPrice(p.SetPrice(1000), p.Price())

	fmt.Printf("Product 구조체 값 : %#v\n", p)

	p2, _ := product.NewProduct(1000, "cola")
	fmt.Printf("NewProduct : %#v\n", p2)
}
