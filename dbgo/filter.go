package dbgo

import (
	"fmt"
	"log"

	"go.etcd.io/bbolt"
)

const (
	FilterTypeEQ = "eq"
)

// func eq(a, b any) bool {
// 	return a == b
// }

type comparison func(a, b any) bool

type compFilter struct {
	kvs  Map //Key-Value Map
	comp comparison
}

func (f compFilter) apply(record Map) bool {
	for k, v := range f.kvs {
		value, ok := record[k]
		if !ok {
			return false // If key is missing in record
		}
		if k == "id" {
			return f.comp(value, uint64(v.(int)))
		}
		return f.comp(value, v)
	}
	return true
}

type Filter struct {
	dbgo        *Dbgo
	coll        string
	compFilters []compFilter
	Slct        []string
	// limit       int
}

func NewFilters(db *Dbgo, coll string) *Filter {
	return &Filter{
		dbgo:        db,
		coll:        coll,
		compFilters: make([]compFilter, 0),
	}
}

func (f *Filter) Insert(values Map) (uint64, error) {

	tx, err := f.dbgo.db.Begin(true)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	collBucket, err := tx.CreateBucketIfNotExists([]byte(f.coll))
	if err != nil {
		return 0, err
	}
	id, err := collBucket.NextSequence()
	if err != nil {
		return 0, err
	}
	b, err := f.dbgo.Encoder.Encode(values)
	if err != nil {
		return 0, err
	}
	if err := collBucket.Put(uint64toBytes(id), b); err != nil {
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return 0, err
	}
	return id, nil

}

func (f *Filter) Find() ([]Map, error) {
	tx, err := f.dbgo.db.Begin(true)
	if err != nil {
		return nil, err
	}
	bucket := tx.Bucket([]byte(f.coll))
	if bucket == nil {
		return nil, fmt.Errorf("bucket [%s] is not found", f.coll)
	}
	records, err := f.findin(bucket)
	if err != nil {
		return nil, err
	}
	fmt.Println("records", records)
	return records, tx.Commit()
}

// func (f *Filter) Update()

// func (f *Filter) Delete()

//	func (f *Filter) Select(values ...string) *Filter {
//		f.slct = append(f.slct, values...)
//		return f
//	}
func (f *Filter) findin(b *bbolt.Bucket) ([]Map, error) {
	response := []Map{}
	b.ForEach(func(k, v []byte) error {
		record := Map{
			"id": uint64FromBytes(k),
		}
		if err := f.dbgo.Decoder.Decode(v, &response); err != nil {
			return err
		}
		include := true
		for _, filter := range f.compFilters {
			if !filter.apply(record) {
				include = false
				break
			}
		}
		if !include {
			return nil
		}
		record = f.applySelect(record)
		response = append(response, record)
		return nil

	})
	return response, nil
}
func (f *Filter) applySelect(record Map) Map {
	if len(f.Slct) == 0 {
		return record
	}
	data := Map{}
	for _, key := range f.Slct {
		if _, ok := record[key]; ok {
			data[key] = record[key]
		}
	}
	return data
}

// This is how data is stored in  Dbgo
// | id(key) | value|
// | 1       | `{"name":"Akanksh","age":23}` |
// | 2       | `{"name":"John","age":25}` |
