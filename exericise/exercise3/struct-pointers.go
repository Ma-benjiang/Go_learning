package main

import "fmt"

type Vertex struct {
	X int
	Y int
}
func (V *Vertex) Swap(){
	V.Y,V.X = V.X,V.Y 
}

func main() {
	v := Vertex{1, 2}
	(&v).Swap()
	fmt.Println(v)
	p := &v
	p.X = 19
	fmt.Println(v)

}