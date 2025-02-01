package main

import (
	"log"

	"github.com/akankshrv/DBgo/dbgo"
)

func main() {

	db, err := dbgo.New()
	if err != nil {
		log.Fatal(err)
	}
	// user := dbgo.Map{
	// 	"name": "kiran",
	// 	"age":  3,
	// }

	filter := dbgo.NewFilters(db, "testcoll")
	// id, err := filter.Insert(user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", id)
	// results, err := filter.Eq(dbgo.Map{"age": float64(30)}).Select("name").Find()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	filter.View()

}
