package dbgo

import (
	"testing"
)

func TestCreateIndex(t *testing.T) {
	db, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize database %v", err)
	}

	err = db.CreateCollection("testcoll")
	if err != nil {
		t.Fatalf("Failed to create collection %v", err)

	}
	err = db.CreateIndex(Index{Collection: "testcoll", Field: "age"})
	if err != nil {
		t.Fatalf("Failed to create index: %v", err)
	}
	tx, err := db.db.Begin(false)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer tx.Rollback()
	indexBucketName := "testcoll_index_age"
	indexBucket := tx.Bucket([]byte(indexBucketName))
	if indexBucket == nil {
		t.Fatalf("Index bucket '%s' does not exist", indexBucketName)
	}
}

func TestUpdateIndex(t *testing.T) {
	db, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize database %v", err)
	}

	if err = db.CreateCollection("testcoll"); err != nil {
		t.Fatalf("Failed to create collection: %v", err)
	}
	if err = db.CreateIndex(Index{Collection: "testcoll", Field: "age"}); err != nil {
		t.Fatalf("Failed to create collection: %v", err)
	}

	record := Map{"id": float64(1), "name": "Ak", "age": 25}
	if err = db.UpdateIndex(Index{Collection: "testcoll", Field: "age"}, record); err != nil {
		t.Fatalf("Failed to update index: %v", err)
	}

	tx, err := db.db.Begin(false)
	if err != nil {
		t.Fatalf("Error in starting the transaction:%v", err)
	}

	defer tx.Rollback()
	indexBucketName := "testcoll_index_age"
	indexBucket := tx.Bucket([]byte(indexBucketName))
	if indexBucket == nil {
		t.Fatalf("Index bucket '%s' does not exist", indexBucketName)
	}

	key := []byte("25")
	value := indexBucket.Get(key)
	if value == nil {
		t.Fatalf("Index entry for key '25' does not exist")
	}

	recordID := uint64FromBytes(value)
	if recordID != 1 {
		t.Errorf("Expected record ID 1, got %d", recordID)
	}
}
