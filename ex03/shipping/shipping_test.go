package shipping

import (
	"ex03/coubang"
	"ex03/dj"
	"ex03/product"
)

func ExampleShipping_coubang_rocket() {
	p, _ := product.NewProduct(200000, "ps4")
	test, _ :=
		coubang.NewCoubang(*p, "yongckim", "minjo", true)
	Shipping(test)
	//Output:상품명 : ps4
	//상품가격 : 200000원
	//[COUBANG] yongckim 님이 minjo 님에게 로켓 배송을 시작했습니다.
}

func ExampleShipping_coubang_send() {
	p, _ := product.NewProduct(200000, "ps4")
	test, _ :=
		coubang.NewCoubang(*p, "yongckim", "minjo", false)
	Shipping(test)
	//Output:상품명 : ps4
	//상품가격 : 200000원
	//[COUBANG] yongckim 님이 minjo 님에게 일반 배송을 시작했습니다.
}

func ExampleShipping_dj() {
	p, err := product.NewProduct(200000, "ps4")
	if err != nil {
		panic(err)
	}
	test, err := dj.NewDj(*p, "yongckim", "jseo")
	if err != nil {
		panic(err)
	}
	Shipping(test)
	//Output: 상품명 : ps4
	//상품가격 : 200000원
	//[DJ] yongckim 님이 jseo 님에게 일반 배송을 시작했습니다.
}
