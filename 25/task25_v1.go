package main

import (
	"fmt"
	"time"
)

func mySleep(sec int) {
	ch := make(chan struct{})
	//используем горутину которая будет проверять просшло время задержки или нет
	go func() {
		s1 := time.Now()
		for {
			if (s1.Add(time.Second * time.Duration(sec))).Before(time.Now()) {
				close(ch) // "пробуждаемся"
				return
			}
		}
	}()
	<-ch //главная горутина блокируется, пока не закроется канал
}

func main() {
	fmt.Println("start")
	mySleep(10)
	fmt.Println("stop")
}
