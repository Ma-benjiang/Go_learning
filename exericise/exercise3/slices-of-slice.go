package main

import "fmt"

func main() {

	var board=make([][5]string,1,5)
	for i :=range board {
		for j :=range board[i] {
			board[i][j]= "x"
		}
	}
	fmt.Println(board)
}