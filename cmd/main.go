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
	user := map[string]interface{}{
		"name": "Akanksh",
		"age":  22,
	}
	filter := dbgo.NewFilters(db, "testcoll")
	id, err := filter.Insert(user)
	if err != nil {
		log.Fatal(err)
	}
	// coll, err := db.CreateCollectio("users")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Printf("%+v\n", id)

}
