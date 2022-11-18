package main

import "fmt"
//引用类型的别名类型的值方法对值的改变也是有效的
type MyIntSlice []int
func (M MyIntSlice) Max()(max int){
	max =0
	for i :=range M{
		if M[i]>max {
			max = M[i]
		}
	}
	for i :=range M{
		if M[i]==max {
			M[i]=1000
		}
	}
	return max
}

func main() {
	primes := MyIntSlice{1, 2, 3, 4, 5, 6}
	fmt.Println(primes[:3].Max())
	fmt.Println(primes)
	// primes.Max()
	fmt.Print(primes.Max())
	fmt.Println(primes)
	
	// fmt.Println()
}
