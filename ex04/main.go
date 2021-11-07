package main

import (
	"ex04/check"
	"ex04/coubang"
	"ex04/product"
	"fmt"
)

func main() {
	p, err := product.NewProduct(200000, "ps4")
	if err != nil {
		fmt.Println(err)
	}
	d, _ := coubang.NewCoubang(*p, "yongckim", "min-jo", false)
	check.Check(d)
}
