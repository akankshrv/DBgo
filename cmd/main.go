package main

import (
	"fmt"
	"log"

	"github.com/akankshrv/DBgo/dbgo"
)

func main() {

	// db, err := dbgo.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // user := dbgo.Map{
	// // 	"name": "kiran",
	// // 	"age":  3,
	// // }

	// filter := dbgo.NewFilters(db, "testcoll")
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
	// err = filter.Eq(dbgo.Map{"age": float64(19)}).Delete()
	// if err != nil {
	// 	log.Fatalf("Failed to delete records: %v", err)
	// } else {
	// 	fmt.Println("Records successfully deleted.")
	// }
	//filter.View()

	db, err := dbgo.New()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	err = db.CreateCollection("testcoll")
	if err != nil {
		log.Fatalf("Failed to create collection: %v", err)
	}

	err = db.CreateIndex(dbgo.Index{Collection: "testcoll", Field: "age"})
	if err != nil {
		log.Fatalf("Failed to create index on 'age': %v", err)
	}

	filter := dbgo.NewFilters(db, "testcoll")
	users := []dbgo.Map{
		{"id": float64(1), "name": "Alice", "age": 25},
		{"id": float64(2), "name": "Bob", "age": 30},
	}

	for _, user := range users {
		id, err := filter.Insert(user)
		if err != nil {
			log.Fatalf("Failed to insert record: %v", err)
		}
		fmt.Printf("Inserted record with ID: %d\n", id)

		err = db.UpdateIndex(dbgo.Index{Collection: "testcoll", Field: "age"}, user)
		if err != nil {
			log.Fatalf("Failed to update index for record: %v", err)
		}
		fmt.Printf("Index updated for record: %+v\n", user)
	}

}
