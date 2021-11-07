package main

import (
	"ex02/coubang"
	"fmt"
)

func main() {
	//p, err := product.NewProduct(1000, "candy")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//d, err := coubang.NewCoubang(*p, "yongckim", "min-jo", false)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//d.Info()
	c := coubang.Coubang{}
	err := c.SetName("candy")
	if err != nil {
		fmt.Println(err)
	}
	err = c.SetPrice(1000)
	if err != nil {
		fmt.Println(err)
	}
	err = c.SetTo("yongckim")
	if err != nil {
		fmt.Println(err)
	}
	err = c.SetFrom("min-jo")
	if err != nil {
		fmt.Println(err)
	}
	c.Info()
	//c.Send()
	//c.RocketSend()

}
