package check

import (
	"fmt"

	"ex04/coubang"
	"ex04/dj"
	"ex04/shipping"
)

func Check(sender shipping.Sender) {
	switch s := sender.(type) {
	case *coubang.Coubang:
		s.Info()
		fmt.Printf("COUBANG 물품입니다.\n")
	case *dj.Dj:
		s.Info()
		fmt.Printf("DJ 배송 물품입니다.\n")
	default:
		fmt.Printf("알 수 없는 배송업체입니다.\n")
	}
}
