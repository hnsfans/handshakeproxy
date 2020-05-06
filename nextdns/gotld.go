package nextdns

import (
	"log"
	"strings"
)

var (
	tldMap = map[string]map[string]bool{}
)

func init() {
	items := strings.Split(nowTLD, "\n")
	for _, v := range items {
		firstC := string(v[0])
		if _, exist := tldMap[firstC]; !exist {
			tldMap[firstC] = map[string]bool{}
		}
		tldMap[firstC][v] = false
	}
}

func isDomainHandshake(domain string) bool {
	splits := strings.Split(domain, ".")
	l := len(splits)
	tld := splits[l-1]
	tldUpper := strings.ToUpper(tld)
	firstC := string(tldUpper[0])
	log.Println(firstC)
	if n, exist := tldMap[firstC]; exist {
		return n[tldUpper]
	}
	return true
}
