package main

import (
	tm "github.com/buger/goterm"
	"time"
)

var doorsChan = make(chan int, 10)
var tyresChan = make(chan int, 10)
var chassisChan = make(chan int, 10)

func tyres(ch chan int) {
	n, i := 0, 0
	for i < 100 {

		time.Sleep(1 * time.Second)
		i = i + 4
		if i%4 == 0 {
			ch <- 4
		} else {
			ch <- i
		}
		tm.MoveCursor(5, 15)
		tm.Println(i)
		tm.Flush()
		n = n + 1
	}
	close(ch)
}

func doors(ch chan int) {
	j, n := 0, 0

	for j < 100 {
		time.Sleep(2 * time.Second)
		j = j + 4
		if j%4 == 0 {
			ch <- 4
		} else {
			ch <- j
		}
		tm.MoveCursor(5, 70)
		tm.Println(j)
		tm.Flush()
		n = n + 1
	}
	close(ch)
}

func chassis(ch chan int) {

	k, n := 0, 0
	for k < 100 {

		time.Sleep(3 * time.Second)
		k = k + 1
		ch <- 1
		tm.MoveCursor(5, 125)
		tm.Println(k)
		tm.Flush()

		n = n + 1
	}
	close(ch)
}

func main() {
	tm.Clear()
	tm.MoveCursor(1, 5)
	tm.Println("~~~TYRES MANUFACTURED~~~")
	tm.MoveCursor(1, 60)
	tm.Println("~~~DOORS MANUFACTURED~~~")
	tm.MoveCursor(1, 110)
	tm.Println("~~~CHASSIS MANUFACTURED~~~")
	tm.Flush()

	go tyres(tyresChan)
	go doors(doorsChan)
	go chassis(chassisChan)

	number := 0

	for y3 := range chassisChan {

		y1 := <-doorsChan
		y2 := <-tyresChan

		if y1 == 4 && y2 == 4 && y3 == 1 {
			number = number + 1
			tm.MoveCursor(20, 50)
			tm.Println(number, " - CAR MADE")
			tm.Flush()
		}

	}

}
