package main

import "fmt"
type MyArray [5]int

func (A MyArray) Max() (max int){
	max = 0
	for i:=range A{
		if A[i]>max {
			max = A[i]
		}
	}
	return max
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0],a[1])
	p := &a
	p[0]="hello"
	fmt.Println(a)
	primes := [6]int{2,3,4,5,6,7}
	fmt.Println(primes)
}