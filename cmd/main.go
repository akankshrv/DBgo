package main

import (
	"fmt"
	"log"

	"github.com/akankshrv/DBgo/dbgo"
)

func main() {

	db, err := dbgo.New()
	if err != nil {
		log.Fatal(err)
	}
	user := map[string]string{
		"name": "Akanksh",
		"age":  "22",
	}
	id, err := db.Insert("users", user)
	if err != nil {
		log.Fatal(err)
	}
	// coll, err := db.CreateCollectio("users")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Printf("%+v\n", id)

}
