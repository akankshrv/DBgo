package dbgo

import (
	"fmt"
	"os"

	"go.etcd.io/bbolt"
)

const (
	defaultDBName = "default"
	extension     = "dbgo"
)

type Map map[string]any

type Dbgo struct {
	currentDatabase string
	*Options
	db *bbolt.DB
}

func New(options ...Optfunc) (*Dbgo, error) {
	opts := &Options{
		Encoder: JSONEncoder{},
		Decoder: JSONDecoder{},
		DBName:  defaultDBName,
	}
	for _, fn := range options {
		fn(opts)
	}

	dbname := fmt.Sprintf("%s.%s", opts.DBName, extension)
	db, err := bbolt.Open(dbname, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &Dbgo{
		currentDatabase: dbname,
		db:              db,
		Options:         opts,
	}, nil
}

func (h *Dbgo) DropDatabase(name string) error {
	dbname := fmt.Sprintf("%s.%s", name, extension)
	return os.Remove(dbname)
}
func (h *Dbgo) CreateCollection(name string, indexes ...string) error {
	tx, err := h.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.CreateBucketIfNotExists([]byte(name))
	if err != nil {
		return err
	}

	return tx.Commit()
}
