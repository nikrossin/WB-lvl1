package main

import "fmt"

//интерфейс для чтения данных
type TimeSecReader interface {
	ReadSecTime()
}

//структура времени в секундах
type TimeSeconds struct {
	dataTimeInSec int
}

func NewTimeSeconds(sec int) *TimeSeconds {
	return &TimeSeconds{sec}
}

//структура времени в сек реализует интерфейс TimeSecReader
func (time *TimeSeconds) ReadSecTime() {
	fmt.Println(time.dataTimeInSec)
}

//структура полного времени
type FullTime struct {
	hours   int
	minutes int
	seconds int
}

func NewFullTime(h int, m int, s int) *FullTime {
	return &FullTime{h, m, s}
}

//структура полного времени имеет метод конвертации в секунды
func (time *FullTime) convertFullToSec() int {
	return time.hours*3600 + time.minutes*60 + time.seconds
}

//адаптер для чтения данных из Полного времени
type AdapterTime struct {
	fullTime *FullTime
}

func NewAdapter(obj *FullTime) *AdapterTime {
	return &AdapterTime{obj}
}

//для адаптера реализуем метод чтения данных
func (time *AdapterTime) ReadSecTime() {
	fmt.Println(time.fullTime.convertFullToSec())
}

func main() {
	fullTime := NewFullTime(1, 34, 40)
	secTime := NewTimeSeconds(3600)
	adapter := NewAdapter(fullTime) //адаптер реализует интерфейс TimeSecReader
	//тем самым не меняя код структуры FullTime, мы можем читать данные засчет адаптера
	sTime := TimeSecReader(secTime)
	fTime := TimeSecReader(adapter)
	sTime.ReadSecTime()
	fTime.ReadSecTime()

}
