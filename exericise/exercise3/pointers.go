package main

import "fmt"

func main() {
	i, j := 42, 2030
	p := &i
	fmt.Println(*p)//read i through the pointer
	*p = 21 //set i through the pointer
	fmt.Println(i)
	p = &j
	*p = *p /37
	fmt.Println(j)
}