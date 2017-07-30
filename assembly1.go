package main

import (
	tm "github.com/buger/goterm"
	"time"
)

var doorsChan = make(chan int, 10)
var tyresChan = make(chan int, 10)
var chassisChan = make(chan int, 10)

func tyres(ch chan int) {
	 tyres :=  0
	for tyres < 100 {

		time.Sleep(1 * time.Second)
		tyres = tyres + 4
		if tyres%4 == 0 {
			ch <- 4
		} else {
			ch <- tyres
		}
		tm.MoveCursor(5, 15)
		tm.Println(tyres)
		tm.Flush()
		
	}
	close(ch)
}

func doors(ch chan int) {
	doors := 0

	for doors < 100 {
		time.Sleep(2 * time.Second)
		doors = doors + 4
		if doors%4 == 0 {
			ch <- 4
		} else {
			ch <- doors
		}
		tm.MoveCursor(5, 70)
		tm.Println(doors)
		tm.Flush()
	
	}
	close(ch)
}

func chassis(ch chan int) {

	chassis :=  0
	for chassis < 100 {

		time.Sleep(3 * time.Second)
		chassis = chassis + 1
		ch <- 1
		tm.MoveCursor(5, 125)
		tm.Println(chassis)
		tm.Flush()
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

	carsNumber := 0

	for y3 := range chassisChan {

		y1 := <-doorsChan
		y2 := <-tyresChan

		if y1 == 4 && y2 == 4 && y3 == 1 {
			carsNumber = carsNumber + 1
			tm.MoveCursor(20, 50)
			tm.Println(carsNumber, " - CAR MADE")
			tm.Flush()
		}

	}

}