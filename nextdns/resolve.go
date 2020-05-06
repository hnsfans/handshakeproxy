package nextdns

import (
	"encoding/json"
	"handshakeproxy/cache"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	client = http.Client{
		Timeout: 3 * time.Minute,
	}
	nextDNSPrefix = "https://dns.nextdns.io/"
	nextDNSID     = "c3484e"
	cacheTTL      = 60
)

func makeResolveURL(domain, resolveType string) string {
	return nextDNSPrefix + nextDNSID + "?name=" + domain + "&type=" + resolveType
}

func resolveFromCache(domain string) string {
	if n, ok := cache.GetKey(domain); ok {
		return n
	}
	return ""
}

func writeCache(domain, value string) bool {
	cache.WriteKV(domain, value, cacheTTL)
	return true
}

// HandshakeResolve 解析 handshake 域名
func HandshakeResolve(domain string) string {
	if n := resolveFromCache(domain); n != "" {
		return n
	}
	var (
		result map[string]interface{}
		err    error
		resp   *http.Response
	)

	if resp, err = client.Get(makeResolveURL(domain, "A")); err == nil {
		defer resp.Body.Close()
		if bodyBytes, err := ioutil.ReadAll(resp.Body); err == nil {
			if err = json.Unmarshal(bodyBytes, &result); err == nil {
				if ans, exist := result["Answer"]; exist {
					l := ans.([]interface{})
					if len(l) != 0 {
						i := l[0].(map[string]interface{})
						d := i["data"].(string)
						writeCache(domain, d)
						return d
					}
				}
			}
		}
	}
	return ""
}
