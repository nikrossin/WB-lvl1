package main

import (
	"fmt"
	"time"
)

func mySleep(sec int) {
	ch := make(chan struct{})
	go func() {
		s1 := time.Now()
		for {
			if (s1.Add(time.Second * time.Duration(sec))).Before(time.Now()) {
				close(ch)
				return
			}
		}
	}()
	<-ch
}

func main() {
	fmt.Println("start")
	mySleep(10)
	fmt.Println("stop")
}
