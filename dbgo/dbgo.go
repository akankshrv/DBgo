package dbgo

import (
	"fmt"

	"go.etcd.io/bbolt"
)

const (
	defaultDBName = "default"
)

type Collection struct {
	bucket *bbolt.Bucket
}
type Dbgo struct {
	db *bbolt.DB
}

func New() (*Dbgo, error) {
	dbname := fmt.Sprintf("%s.Dbgo", defaultDBName)
	db, err := bbolt.Open(dbname, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &Dbgo{
		db: db,
	}, nil
}

func (h *Dbgo) CreateCollectio(name string) (*Collection, error) {
	coll := Collection{}
	err := h.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte(name))
		if err != nil {
			return err
		}
		coll.bucket = bucket
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &coll, nil
}

// db.Update(func(tx *bbolt.Tx) error {

// 		id := uuid.New()
// 		for k, v := range user {
// 			if err := bucket.Put([]byte(k), []byte(v)); err != nil {
// 				return err
// 			}
// 		}
// 		if err := bucket.Put([]byte("id"), []byte(id.String())); err != nil {
// 			return err
// 		}

// 		return nil
// 	})
