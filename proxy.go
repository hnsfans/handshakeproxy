package main

import (
	"handshakeproxy/cache"
	"handshakeproxy/nextdns"
	"net/http"

	"github.com/elazarl/goproxy"
	log "github.com/sirupsen/logrus"
)

func resoveHostDomain(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	nHost := nextdns.HandshakeResolve(req.URL.Host)
	if nHost != "" {
		req.URL.Host = nHost
	}
	return req, nil
}

func main() {
	cache.Info()
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true})
	log.Println("start handshake proxy...")
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(resoveHostDomain)
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
