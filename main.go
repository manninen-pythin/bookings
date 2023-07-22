package main

import (
	"encoding/json"
	"fmt"
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

	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},

		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	var unmarshalled []helpers.Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	log.Printf("unmarshalled: %v", unmarshalled)

	// write json from a struct
	var mySlice []helpers.Person

	var m1 helpers.Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	var m2 helpers.Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m1, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "     ")

	if err != nil {
		log.Println("error:", err)
	}

	fmt.Println(string(newJson))

}
