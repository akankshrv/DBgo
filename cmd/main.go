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
	// results, err := filter.Lt(dbgo.Map{"age": float64(20)}).Select("name").Find()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// results, err := filter.Eq(dbgo.Map{"age": float64(13)}).Update(dbgo.Map{"age": 19})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(results)
	err = filter.Eq(dbgo.Map{"age": float64(19)}).Delete()
	if err != nil {
		log.Fatalf("Failed to delete records: %v", err)
	} else {
		fmt.Println("Records successfully deleted.")
	}
	//filter.View()
}
