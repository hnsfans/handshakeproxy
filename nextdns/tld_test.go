package nextdns

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSplitTld(t *testing.T) {
	items := strings.Split(nowTLD, "\n")
	result := map[string]map[string]bool{}
	for _, v := range items {
		x := len(v)
		lastC := v[x-1:]
		if _, exist := result[lastC]; !exist {
			result[lastC] = map[string]bool{}
		}
		result[lastC][v] = true
	}

	jsonBytes, _ := json.Marshal(result)
	t.Log(string(jsonBytes))

	t.Errorf("")
}

func TestIsHandShake(t *testing.T) {
	t.Log(isDomainHandshake("welcome.2d"))
	t.Log(isDomainHandshake("baidu.com"))
	t.Log(isDomainHandshake("welcome.nb"))
	t.Errorf("")
}
