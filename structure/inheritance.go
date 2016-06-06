package main

import "fmt"

type A interface {
	A() B
}

type AImpl struct {
}

func (a AImpl) A() C {
	ret
}

type B interface {
	B() bool
}

func main() {
	a := A{true}
	fmt.Printf("A: %+v\n", a)

	b := B{1}
	fmt.Printf("B: %+v\n", b)

	simple := Simple{a: A{true}, b: B{1}}
	fmt.Printf("Simple: %+v\n", simple)

	fused := Fused{A: A{true}, B: B{1}}
	fmt.Printf("Fused: %+v\n", fused)
}
