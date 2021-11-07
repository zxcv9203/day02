package coubang

import (
	"reflect"
	"testing"

	"ex04/product"
)

func TestCoubang_From(t *testing.T) {
	coubang := Coubang{}
	_ = coubang.SetFrom("yongckim")
	got := coubang.From()
	want := "yongckim"
	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}

func TestCoubang_To(t *testing.T) {
	coubang := Coubang{}
	_ = coubang.SetFrom("yongckim")
	got := coubang.From()
	want := "yongckim"
	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}

func ExampleCoubang_Info_noName() {
	coubang := Coubang{}
	coubang.Info()
	//Output: 상품명을 먼저 설정해주세요
}

func ExampleCoubang_Info_noPrice() {
	p := product.Product{}
	_ = p.SetName("candy")
	coubang := Coubang{p, "", "", false}
	coubang.Info()
	//Output: 상품의 가격을 먼저 설정해주세요

}

func ExampleCoubang_Info() {
	p, err := product.NewProduct(1000, "candy")
	if err != nil {
		panic(err)
	}
	coubang := Coubang{*p, "", "", false}
	coubang.Info()
	//Output: 상품명 : candy
	//상품가격 : 1000원
}

func ExampleCoubang_Send_info_noName() {
	coubang := Coubang{}
	coubang.Send()
	//Output: 상품명을 먼저 설정해주세요
}

func ExampleCoubang_Send_info_noPrice() {
	p := product.Product{}
	_ = p.SetName("candy")
	coubang := Coubang{p, "", "", false}
	coubang.Send()
	//Output: 상품의 가격을 먼저 설정해주세요
}

func ExampleCoubang_Send_noTo() {
	p, _ := product.NewProduct(1000, "candy")
	coubang := Coubang{*p, "", "", false}
	coubang.Send()
	//Output: 상품명 : candy
	//상품가격 : 1000원
	//택배 발송 전 발송자를 먼저 입력해주세요
}

func ExampleCoubang_Send_noFrom() {
	p, _ := product.NewProduct(1000, "candy")
	coubang := Coubang{*p, "yongckim", "", false}
	coubang.Send()
	//Output: 상품명 : candy
	//상품가격 : 1000원
	//택배 발송 전 수신자를 먼저 입력해주세요
}

func ExampleCoubang_Send() {
	p, _ := product.NewProduct(1000, "candy")
	coubang := Coubang{*p, "yongckim", "min-jo", false}
	coubang.Send()
	//Output: 상품명 : candy
	//상품가격 : 1000원
	//[COUBANG] yongckim 님이 min-jo 님에게 일반 배송을 시작했습니다.
}

func ExampleCoubang_RocketSend() {
	p, _ := product.NewProduct(1000, "candy")
	coubang := Coubang{*p, "yongckim", "jseo", true}
	coubang.RocketSend()
	//Output: 상품명 : candy
	//상품가격 : 1000원
	//[COUBANG] yongckim 님이 jseo 님에게 로켓 배송을 시작했습니다.
}

func TestCoubang_NewCoubang(t *testing.T) {
	p, _ := product.NewProduct(1000, "candy")
	got, _ := NewCoubang(*p, "yongckim", "jseo", true)
	want := &Coubang{*p, "yongckim", "jseo", true}
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("expected %#v but got %#v", want, got)
	}
}
