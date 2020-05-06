package cache

import (
	"encoding/json"
	"fmt"
	"time"
)

// ExpireData 会过期的数据对象
type ExpireData struct {
	Value  string `json:"value"`
	Expire int64  `json:"expire"`
}

// GetExpire 通过key 获取数据对象，
// 该数据对象是一个json 字符串，
// data 为数据本身
// expire 为过期的时间戳，
// 如果该数据已经过期，那么从数据库中，将其删除
func GetExpire(key string) (string, bool) {
	v, exist := get(key)
	if exist {
		n := time.Now().Unix()
		item := &ExpireData{}
		json.Unmarshal([]byte(v), &item)
		// 数据过期了
		if n >= item.Expire {
			remove(key)
			return "", false
		}
		return item.Value, true
	}
	return "", false
}

// SetExpire 设置数据，包含过期时间，
func SetExpire(key string, value string, ttl int64) {
	itemData := ExpireData{
		Value:  value,
		Expire: time.Now().Unix() + ttl,
	}

	realBytes, err := json.Marshal(itemData)
	if err == nil {
		fmt.Println("set ", key, string(realBytes))
		set(key, string(realBytes))
	} else {
		fmt.Println("json encode error .", err)
	}
}
