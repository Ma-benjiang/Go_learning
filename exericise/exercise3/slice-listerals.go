package main

import (
	"fmt"
)

func main(){
	q :=[]int{2,3,4,5,6,6}
	q = q[1:4]
	fmt.Println(cap(q))

	s :=[]struct{
		i int 
		b bool
	}{
		{2,true},
		{3,false},
		{5,true},
		{11,false},
	}
	fmt.Println(s)
}