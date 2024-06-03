package logic

import (
	"fmt"
	"testing"
)

type Defaulter interface {
	InitDefault()
}

type Foo struct {
	Message string
	Value   int64
}

func (f *Foo) InitDefault() {
	f.Message = "Foo 默认值"
	f.Value = 100
}

type Bar struct {
	Message string
	Value   float64
}

func (b *Bar) InitDefault() {
	b.Message = "Bar 默认值"
	b.Value = 100.5
}

type defaulterPtr[T any] interface {
	*T
	Defaulter
}

func Default[T any, P defaulterPtr[T]]() T {
	var v T
	P.InitDefault(&v)
	return v
}

func TestInitDefault(t *testing.T) {
	foo := Default[Foo]()
	fmt.Println(foo)
	bar := Default[Bar]()
	fmt.Println(bar)
}
