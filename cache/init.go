package cache

import (
	"log"

	"github.com/boltdb/bolt"
)

var (
	db         *bolt.DB
	err        error
	bucketName = "handshake"
)

// Init 数据库初始化函数
func Init() {
	db, err = bolt.Open("cache.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(bucketName))
		return nil
	})
}
