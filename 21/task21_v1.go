package main

import "fmt"

type TimeReader interface {
	ReadTime()
}

type TimeSeconds struct {
	dataTimeInSec int
}

func NewTimeSeconds(sec int) *TimeSeconds {
	return &TimeSeconds{sec}
}

func (time *TimeSeconds) ReadTime() {
	fmt.Println(time.dataTimeInSec)
}

type FullTime struct {
	hours   int
	minutes int
	seconds int
}

func NewFullTime(h int, m int, s int) *FullTime {
	return &FullTime{h, m, s}
}

func (time *FullTime) convertFullToSec() int {
	return time.hours*3600 + time.minutes*60 + time.seconds
}

type AdapterTime struct {
	fullTime *FullTime
}

func NewAdapter(obj *FullTime) *AdapterTime {
	return &AdapterTime{obj}
}

func (time *AdapterTime) ReadTime() {
	fmt.Println(time.fullTime.convertFullToSec())
}

func main() {
	fullTime := NewFullTime(1, 34, 40)
	secTime := NewTimeSeconds(3600)
	adapter := NewAdapter(fullTime)

	sTime := TimeReader(secTime)
	fTime := TimeReader(adapter)
	sTime.ReadTime()
	fTime.ReadTime()

}
