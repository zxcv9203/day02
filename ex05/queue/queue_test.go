package queue

import (
	"ex05/coubang"
	"ex05/product"
	"testing"
)

func TestQueue_Empty(t *testing.T) {
	t.Run("Queue_Empty", func(t *testing.T) {
		q := Queue{}
		got := q.Empty()
		want := true
		if got != want {
			t.Errorf("expected %#v but got %#v", want, got)
		}
	})
	t.Run("Queue_NotEmpty", func(t *testing.T) {
		q := Queue{}
		p, err := product.NewProduct(1000, "candy")
		if err != nil {
			t.Error(err)
		}
		c, err := coubang.NewCoubang(*p, "yongckim", "min-jo", false)
		if err != nil {
			t.Error(err)
		}
		q.Push(c)
		got := q.Empty()
		want := false
		if got != want {
			t.Errorf("expected %#v but got %#v", want, got)
		}
	})
}

func TestQueue_Size(t *testing.T) {
	q := Queue{}
	p, err := product.NewProduct(1000, "candy")
	if err != nil {
		t.Error(err)
	}
	c, err := coubang.NewCoubang(*p, "yongckim", "min-jo", false)
	if err != nil {
		t.Error(err)
	}
	q.Push(c)
	got := q.Size()
	want := 1
	if got != want {
		t.Errorf("expected %#v but got %#v", want, got)
	}
}

func TestQueue_Pop(t *testing.T) {
	q := Queue{}
	p, err := product.NewProduct(1000, "candy")
	if err != nil {
		t.Error(err)
	}
	c, err := coubang.NewCoubang(*p, "yongckim", "min-jo", false)
	if err != nil {
		t.Error(err)
	}
	q.Push(c)
	got := q.Pop()
	want := c
	if got != want {
		t.Errorf("expected %#v but got %#v", want, got)
	}
}

// 나중에 사용자가 만든 Push 하고 비교하게끔 테스트 케이스를 만들어야함
func TestQueue_Push(t *testing.T) {
	q := Queue{}
	p, err := product.NewProduct(1000, "candy")
	if err != nil {
		t.Error(err)
	}
	c, err := coubang.NewCoubang(*p, "yongckim", "min-jo", false)
	if err != nil {
		t.Error(err)
	}
	q.Push(c)
	got := q.Size()
	want := 1
	if got != want {
		t.Errorf("expected %#v but got %#v", want, got)
	}
}
