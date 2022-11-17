package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// 在log信息前面打印greetings,不带有时间戳，源文件信息
	log.SetPrefix("greetings:")
	log.SetFlags(0)
	// message,err:= greetings.Hello("ma")
	names := []string{"me","ben","jiang"}
	// 打印错误，停止程序
	messages,err :=greetings.Hellos(names)
	if err !=nil{
		log.Fatal(err)
	}
	// 没有错误打印消息
	fmt.Println(messages)
}