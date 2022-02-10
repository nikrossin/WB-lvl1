package main

import (
	"fmt"
	"time"
)

//используем готовую функцию, которая возвращает канал
func mySleep(sec int) {
	<-time.After(time.Duration(sec) * time.Second)

}
func main() {
	fmt.Println("start")
	mySleep(5)
	fmt.Println("stop")
}
