package dbgo

import "fmt"

type Index struct {
	Collection string //Name of the collection
	Field      string //Field to be indexed
}

// Index record map stores index.Field and id (Mapping)
func (h *Dbgo) CreateIndex(index Index) error {
	tx, err := h.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	collBucket := tx.Bucket([]byte(index.Collection))
	if collBucket == nil {
		return fmt.Errorf("collection %s doesnt exist", index.Collection)
	}

	indexBucketName := fmt.Sprintf("%s_index_%s", index.Collection, index.Field)

	_, err = tx.CreateBucketIfNotExists([]byte(indexBucketName))
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (h *Dbgo) UpdateOnIndex(index Index, record Map) error {
	tx, err := h.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	indexBucketName := fmt.Sprintf("%s_index_%s", index.Collection, index.Field)
	indexBucket := tx.Bucket([]byte(indexBucketName))
	if indexBucket == nil {
		return fmt.Errorf("collection %s doesnt exist", index.Collection)

	}

	fieldValue, ok := record[index.Field]
	if !ok {
		return fmt.Errorf("filed %s is not found", index.Field)
	}

	idFloat, ok := record["id"].(float64)
	if !ok {
		return fmt.Errorf("record id is not a valid uint64")
	}
	id := uint64(idFloat)

	k := []byte(fmt.Sprintf("%v", fieldValue))
	idBytes := uint64toBytes(id)

	if err := indexBucket.Put(k, idBytes); err != nil {
		return fmt.Errorf("failed to update index for field '%s': %v", index.Field, err)
	}

	return tx.Commit()

}
