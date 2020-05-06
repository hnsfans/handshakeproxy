package cache

import "log"

// WriteKV 写入 KV 数据
func WriteKV(k, v string, ttl int) {
	log.Println("Write Cache ", k, v, ttl)
	SetExpire(k, v, int64(ttl))
}

// GetKey 读出数据
func GetKey(k string) (string, bool) {
	return GetExpire(k)
}

// Info 简单调用
func Info() {
	log.Println("Start Cache Info Object ..")
}
