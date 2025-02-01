package dbgo

const (
	FilterTypeEQ = "eq"
)

func eq(a, b any) bool {
	return a == b
}

type comparison func(a, b any) bool

type compFilter struct {
	kvs  Map //Key-Value Map
	comp comparison
}

// func ( f compFilter) apply(record Map) bool {
// 	for k, v := range f.kvs {
// 		value, ok := record[k]
// 		if !ok {
// 			return false // If key is missing in record
// 		}
// 		if k == "id"{
// 			 return f.comp(value, uint64(v.(int)))
// 		}
// 		return f.comp(value,v)
// 	}
// 	return true
// }

type Filter struct {
	dbgo        *Dbgo
	coll        string
	compFilters []compFilter
	slct        []string
	limit       int
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
	return id, tx.Commit()

}

// func (h *Dbgo) Find(coll string, k string, query any) (M, error) {
// 	tx, err := h.db.Begin(true)
// 	if err != nil {
// 		return nil, err
// 	}
// 	bucket := tx.Bucket([]byte(coll))
// 	if bucket == nil {
// 		return nil, fmt.Errorf("collection (%s) not found", coll)
// 	}
// 	records, err :=
// }
