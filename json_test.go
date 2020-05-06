package main

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/json"
)

var (
	c = `
{"Status":0,"TC":false,"RD":true,"RA":true,"AD":false,"CD":false,"Question":[{"name":"welcome.2d.","type":1}],"Answer":[{"name":"welcome.2d.","type":5,"TTL":3082,"data":"welcome.2d.s3-website-us-west-1.amazonaws.com."},{"name":"welcome.2d.s3-website-us-west-1.amazonaws.com.","type":5,"TTL":29,"data":"s3-website-us-west-1.amazonaws.com."},{"name":"s3-website-us-west1.amazonaws.com.","type":1,"TTL":30,"data":"52.219.116.10"}],"Additional":[{"name":".","type":41,"TTL":0,"data":"\n;; OPT PSEUDOSECTION:\n; EDNS: version 0; flags: ; udp: 1220"}]}
`
	result map[string]interface{}
)

func TestJson(t *testing.T) {
	err := json.Unmarshal([]byte(c), &result)
	t.Log(err)
	if ans, exist := result["Answer"]; exist {
		l := ans.([]interface{})
		if len(l) != 0 {
			i := l[0].(map[string]interface{})
			t.Log(i["data"])
		}
	}
	t.Error(".")
}
