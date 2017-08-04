//4 tyres manufacture in 1 second
//4 doors manufacture in 3 seconds
//1 chassis manufacture in 2 seconds

package main

import (
	_ "fmt"
	tm "github.com/buger/goterm"
	"sync"
	"time"
)

var doorsChan = make(chan int)
var tyresChan = make(chan int)
var chassisChan = make(chan int)

func tyres(ch chan int) {
	defer wg.Done()

	tyres := 0
	for tyres < 100 {

		time.Sleep(1 * time.Second)
		tm.MoveCursor(35, 65)
		tm.Println("Time : ", time.Now().Format("15:04:05.000000"))
		tyres = tyres + 1
		tm.MoveCursor(5, 15)
		tm.Println(tyres)

		ch <- 1
		tm.Flush()
	}
}

func doors(ch chan int) {
	defer wg.Done()
	doors := 0

	for doors < 100 {
		time.Sleep(3 * time.Second)
		doors = doors + 1
		ch <- 1
		tm.MoveCursor(5, 70)
		tm.Println(doors)
		tm.Flush()

	}
}

func chassis(ch chan int) {
	defer wg.Done()

	chassis := 0
	for chassis < 100 {
		time.Sleep(2 * time.Second)
		chassis = chassis + 1
		ch <- 1
		tm.MoveCursor(5, 125)
		tm.Println(chassis)

		tm.Flush()
	}

}

var wg sync.WaitGroup

func main() {
	tm.Clear()
	tm.MoveCursor(1, 5)
	tm.Println("~~~TYRES MANUFACTURED~~~")
	tm.MoveCursor(1, 60)
	tm.Println("~~~DOORS MANUFACTURED~~~")
	tm.MoveCursor(1, 110)
	tm.Println("~~~CHASSIS MANUFACTURED~~~")
	tm.Flush()

	wg.Add(3)
	go tyres(tyresChan)
	go doors(doorsChan)
	go chassis(chassisChan)

	quit := make(chan int)
	go func() {
		wg.Wait()
		quit <- 1
	}()

	carsNumber := 0
	t, d, c := 0, 0, 0

	go func() {

		for {
			if t > 3 && d > 3 && c > 0 {

				t = t - 4
				d = d - 4
				c = c - 1
				carsNumber = carsNumber + 1

			}
		}
	}()

	for {

		select {
		case <-doorsChan:
			d += 1

		case <-tyresChan:
			t += 1

		case <-chassisChan:
			c += 1

		case <-quit:
			close(tyresChan)
			close(doorsChan)
			close(chassisChan)
			return
		}

		if carsNumber > 1 {

			tm.MoveCursor(20, 30)
			tm.Println(carsNumber, " - CAR MADE")
			tm.Flush()
		}

	}
}
