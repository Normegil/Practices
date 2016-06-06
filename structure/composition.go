package main

import "fmt"

type A struct {
	a bool
}

type B struct {
	b int
}

type Simple struct {
	a A
	b B
}

type Fused struct {
	A
	B
}

type Pointers struct {
	*A
	*B
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
