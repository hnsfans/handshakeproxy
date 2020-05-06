package cache

import "github.com/boltdb/bolt"

func get(key string) (value string, exist bool) {
	var v []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v = b.Get([]byte(key))
		return nil
	})
	if v == nil {
		return "", false
	}
	return string(v), true
}

func set(k, v string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if err := b.Put([]byte(k), []byte(v)); err != nil {
			return err
		}
		return nil
	})
}

func remove(k string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		b.Delete([]byte(k))
		return nil
	})
}
