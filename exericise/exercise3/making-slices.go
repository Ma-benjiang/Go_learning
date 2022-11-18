package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSclice("a",a)

	b :=make([]int,0,5)
	printSclice("b",b)

	c :=b[:2]
	printSclice("c",c)

	d :=c[2:5]
	printSclice("d",d)

}
func printSclice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)

}