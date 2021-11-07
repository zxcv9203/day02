package main

import (
	"ex01/dj"
	"ex01/product"
	"fmt"
)

func main() {
	p, err := product.NewProduct(1000, "candy")
	if err != nil {
		fmt.Println(err)
	}
	d, err := dj.NewDj(*p, "yongckim", "min-jo")
	if err != nil {
		fmt.Println(err)
	}
	d.Info()
	//d2 := dj.Dj{}
	//err := d2.SetName("candy")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = d2.SetPrice(1000)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = d2.SetTo("yongckim")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = d2.SetFrom("jseo")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//d2.Send()
}
