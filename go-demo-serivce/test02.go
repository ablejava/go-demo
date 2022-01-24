package main

import "fmt"

func main()  {
	// int
	var a int = 1
	a = a + 20
	fmt.Println(a)

	b:=1
	c:=3.14
	fmt.Println(b,c)

	d,e,f:=4,5,6
	fmt.Println(d,e,f)

	d,e,f=f,e,d
	fmt.Println(d,e,f)

	fmt.Printf("%d\n",d)
	fmt.Printf("%s\n","hello word")

}
