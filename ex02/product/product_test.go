package product

import (
	"reflect"
	"testing"
)

func TestProduct(t *testing.T) {
	t.Run("price", func(t *testing.T) {
		p := Product{}
		_ = p.SetPrice(1000)
		got := p.Price()
		want := uint(1000)
		if got != want {
			t.Errorf("expected %d but got %d", got, want)
		}
	})
	t.Run("name", func(t *testing.T) {
		p := Product{}
		_ = p.SetName("ps4")
		got := p.Name()
		want := "ps4"
		if got != want {
			t.Errorf("expected %q but got %q", got, want)
		}
	})
	t.Run("constructor", func(t *testing.T) {
		got, _ := NewProduct(1000, "ps4")
		want := &Product{price: 1000, name: "ps4"}
		if reflect.DeepEqual(got, want) == false {
			t.Errorf("expected %#v but got %#v", got, want)
		}
	})
}
