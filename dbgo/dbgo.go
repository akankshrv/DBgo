package dbgo

import (
	"fmt"

	"go.etcd.io/bbolt"
)

const (
	defaultDBName = "default"
)

type M map[string]string

type Collection struct {
	*bbolt.Bucket
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

func (h *Dbgo) CreateCollection(name string) (*Collection, error) {
	tx, err := h.db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	bucket, err := tx.CreateBucketIfNotExists([]byte(name))
	if err != nil {
		return nil, err
	}
	return &Collection{Bucket: bucket}, nil
}

func (h *Dbgo) Insert(collName string, data M) (uint64, error) {

	tx, err := h.db.Begin(true)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	bucket, err := tx.CreateBucketIfNotExists([]byte(collName))
	if err != nil {
		return 0, err
	}
	id, err := bucket.NextSequence()
	if err != nil {
		return 0, err
	}
	for k, v := range data {
		if err := bucket.Put([]byte(k), []byte(v)); err != nil {
			return id, err
		}
	}
	if err := bucket.Put([]byte("id"), uint64toBytes(id)); err != nil {
		return id, err
	}

	return id, tx.Commit()

}

// get http://localhost:4000/users?eq.name={akanksh}
// func (h *Dbgo) Select(coll string, k string, query any) (M, error) {
// 	tx, err := h.db.Begin(false)
// 	if err != nil {
// 		return nil, err
// 	}
// 	bucket := tx.Bucket([]byte(coll))
// 	if bucket == nil {
// 		return nil, fmt.Errorf("collection (%s) not found", coll)
// 	}

// }
