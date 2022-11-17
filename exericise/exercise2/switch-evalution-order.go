package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("when's Saturday")
	today :=time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Tady:")
	case today + 1:
		fmt.Println("Tomorrow")
	default:
		fmt.Println("Too far away")
	}
}