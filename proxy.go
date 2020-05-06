package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/elazarl/goproxy"
	log "github.com/sirupsen/logrus"
)

func nextDNSResolve(domain string) string {
	client := http.Client{
		Timeout: 3 * time.Minute,
	}
	var result map[string]interface{}
	if resp, err := client.Get("https://dns.nextdns.io/c3484e?name=" + domain + "&type=A"); err == nil {
		defer resp.Body.Close()
		if bodyBytes, err := ioutil.ReadAll(resp.Body); err == nil {
			if err := json.Unmarshal(bodyBytes, &result); err == nil {
				if ans, exist := result["Answer"]; exist {
					l := ans.([]interface{})
					if len(l) != 0 {
						i := l[0].(map[string]interface{})
						d := i["data"].(string)
						return d
					}
				}
			}
		}
	} else {
		log.Println(err)
	}

	return "localhost"
}

func resoveHostDomain(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	nHost := nextDNSResolve(req.URL.Host)
	log.Println(nHost)
	if nHost != "localhost" {
		req.URL.Host = nHost
	}
	return req, nil
}

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true})
	log.Println(".")
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(resoveHostDomain)
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
