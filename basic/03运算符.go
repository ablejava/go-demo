package main

import "fmt"

func main() {
	a := 3
	b := 2
	a++
	b--
	c := a/b
	d := a%b
	fmt.Println("a++ =", a)
	fmt.Println("b-- =", b)
	fmt.Println("a/b =", c)
	fmt.Println("a%b =", d)
	a += b
	fmt.Println("a = ", a)
	a -= b
	fmt.Println("a = ", a)
	a *= b
	fmt.Println("a = ", a)
	a /= b
	fmt.Println("a = ", a)

}