package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)


func Hello(name string) (string,error) {
	// 如果名字没给，就返回一个错误
	if name ==""{
		return "",errors.New("empty name")
	}
	// message := fmt.Sprintf("Hi,%v.Welcome!", name)
	message :=fmt.Sprintf(randomFromat(),name)
	return message,nil
}
// 代码包初始化函数在main函数之前只需完成
func init(){
	rand.Seed(time.Now().UnixNano())
}

func Hellos(names []string)(map[string]string,error){
	messages :=make(map[string]string)
	for _ ,name :=range names {
		message,err := Hello(name)
		if err !=nil {
			return nil,err
		}
		messages[name]=message
	}
	return messages,nil
}
func randomFromat() (string) {
	// 消息切片
	formats := []string{
		"Hi,%v.Welcome!",
		"Great to see you,%v!",
		"Hail,%v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
