package nextdns

import (
	"strings"
)

var (
	tldMap = map[string]map[string]int{}
)

func init() {
	items := strings.Split(nowTLD, "\n")
	for _, v := range items {
		firstC := string(v[0])
		if _, exist := tldMap[firstC]; !exist {
			tldMap[firstC] = map[string]int{}
		}
		tldMap[firstC][v] = 1
	}
}

func isDomainHandshake(domain string) bool {
	splits := strings.Split(domain, ".")
	l := len(splits)
	tld := splits[l-1]
	tldUpper := strings.ToUpper(tld)
	firstC := string(tldUpper[0])
	if n, exist := tldMap[firstC]; exist {
		return n[tldUpper] != 1
	}
	return true
}
