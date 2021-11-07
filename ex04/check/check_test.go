package check

import (
	"fmt"

	"ex04/coubang"
	"ex04/dj"
	"ex04/product"
)

func ExampleCheck_coubang() {
	p, _ := product.NewProduct(200000, "ps4")
	test, _ :=
		coubang.NewCoubang(*p, "yongckim", "minjo", false)
	Check(test)
	//Output:상품명 : ps4
	//상품가격 : 200000원
	//COUBANG 물품입니다.
}

func ExampleCheck_dj() {
	p, _ := product.NewProduct(200000, "ps4")
	test, _ := dj.NewDj(*p, "yongckim", "jseo")
	Check(test)
	//Output:상품명 : ps4
	//상품가격 : 200000원
	//DJ 배송 물품입니다.
}

type test struct {
}

func (t *test) Send() {
	fmt.Println("hmm")
}
func ExampleCheck_other() {
	Check(&test{})
	//Output: 알 수 없는 배송업체입니다.
}
