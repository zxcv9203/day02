package dj

import (
	"reflect"
	"testing"

	"ex04/product"
)

func TestDj_From(t *testing.T) {
	dj := Dj{}
	err := dj.SetFrom("yongckim")
	if err != nil {
		panic(err)
	}
	got := dj.From()
	want := "yongckim"
	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}

func TestDj_To(t *testing.T) {
	dj := Dj{}
	err := dj.SetFrom("yongckim")
	if err != nil {
		panic(err)
	}
	got := dj.From()
	want := "yongckim"
	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}

func ExampleDj_Info_noName() {
	dj := Dj{}
	dj.Info()
	//Output: 상품명을 먼저 설정해주세요
}

func ExampleDj_Info_noPrice() {
	p := product.Product{}
	err := p.SetName("candy")
	if err != nil {
		panic(err)
	}
	dj := Dj{p, "", ""}
	dj.Info()
	//Output: 상품의 가격을 먼저 설정해주세요

}

func ExampleDj_Info() {
	p, _ := product.NewProduct(1000, "candy")
	dj := Dj{*p, "", ""}
	dj.Info()
	//Output: 상품명 : candy
	//상품가격 : 1000원
}

func ExampleDj_Send_info_noName() {
	dj := Dj{}
	dj.Send()
	//Output: 상품명을 먼저 설정해주세요
}

func ExampleDj_Send_info_noPrice() {
	p := product.Product{}
	err := p.SetName("candy")
	if err != nil {
		panic(err)
	}
	dj := Dj{p, "", ""}
	dj.Send()
	//Output: 상품의 가격을 먼저 설정해주세요
}

func ExampleDj_Send_noTo() {
	p, err := product.NewProduct(1000, "candy")
	if err != nil {
		panic(err)
	}
	dj := Dj{*p, "", ""}
	dj.Send()
	//Output: 상품명 : candy
	//상품가격 : 1000원
	//택배 발송 전 발송자를 먼저 입력해주세요
}

func ExampleDj_Send_noFrom() {
	p, _ := product.NewProduct(1000, "candy")
	dj := Dj{*p, "yongckim", ""}
	dj.Send()
	//Output: 상품명 : candy
	//상품가격 : 1000원
	//택배 발송 전 수신자를 먼저 입력해주세요
}

func ExampleDj_Send() {
	p, _ := product.NewProduct(1000, "candy")
	dj := Dj{*p, "yongckim", "jseo"}
	dj.Send()
	//Output: 상품명 : candy
	//상품가격 : 1000원
	//[DJ] yongckim 님이 jseo 님에게 일반 배송을 시작했습니다.
}

func TestDj_NewDj(t *testing.T) {
	p, err := product.NewProduct(1000, "candy")
	if err != nil {
		panic(err)
	}
	got, err := NewDj(*p, "yongckim", "min-jo")
	if err != nil {
		panic(err)
	}
	want := &Dj{*p, "yongckim", "min-jo"}
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("expected %#v but got %#v", want, got)
	}

}
