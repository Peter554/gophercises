package main

import (
	"fmt"
	"log"

	"github.com/peter554/gophercises/08-phone-normalizer/db"
)

func main() {
	database, e := db.New()
	if e != nil {
		log.Fatal(e)
	}
	defer database.Close()

	for _, u := range database.GetAllUsers() {
		normalized, e := normalize(u.PhoneNumber)
		if e != nil {
			log.Fatal(e)
		}
		if normalized != u.PhoneNumber {
			u.PhoneNumber = normalized
			database.UpdateUser(u)
		}
	}

	for _, u := range database.GetAllUsers() {
		fmt.Println(u)
	}
}
