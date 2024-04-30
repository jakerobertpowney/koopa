package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
)

func boltDB() (*bolt.DB, error) {

	db, err := bolt.Open("settings.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}
	return db, nil

}

func boltSave(bucket string, key string, i interface{}) error {

	db, err := boltDB()
	if err != nil {
		return fmt.Errorf("error initilasing database driver: %v", err)
	}

	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("error connecting to database bucket: %v", err)
		}

		setting, err := json.Marshal(i)
		if err != nil {
			return fmt.Errorf("error marshalling database record: %v", err)
		}

		if err := b.Put([]byte(key), setting); err != nil {
			return fmt.Errorf("error saving database record: %v", err)
		}

		return nil
	})

}

func boltGet(bucket string, key string, i interface{}) error {

	db, err := boltDB()
	if err != nil {
		return fmt.Errorf("error initilasing database driver: %v", err)
	}

	defer db.Close()

	return db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("no bucket found: %v", bucket)
		}

		value := b.Get([]byte(key))
		if value == nil {
			return fmt.Errorf("no record found: %v", key)
		}

		if err := json.Unmarshal(value, i); err != nil {
			return fmt.Errorf("error unmarshalling database record: %v", err)
		}

		return nil

	})
}

func boltAll(bucket string, i interface{}) error {

	db, err := boltDB()
	if err != nil {
		return fmt.Errorf("error initilasing database driver: %v", err)
	}

	defer db.Close()

	elements := make([][]byte, 0)

	err = db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("no bucket found: %v", bucket)
		}

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			elements = append(elements, v)
		}

		return nil

	})

	if err != nil {
		return fmt.Errorf("no records found: %v", err)
	}

	combinedElements := []byte(fmt.Sprintf("[%s]", bytes.Join(elements, []byte{','})))

	if err := json.Unmarshal(combinedElements, &i); err != nil {
		return fmt.Errorf("error unmarshalling database records: %v", err)
	}

	return nil

}
