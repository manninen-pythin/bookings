package main

import (
	"log"

	"github.com/manninen-pythin/bookings/helpers"
)

const numbPool = 100

func CalcValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numbPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalcValue(intChan)

	num := <-intChan
	log.Println(num)
}
