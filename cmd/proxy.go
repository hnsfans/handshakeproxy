package cmd

import (
	"handshakeproxy/cache"
	"handshakeproxy/nextdns"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/spf13/cobra"
)

func resoveHostDomain(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	nHost := nextdns.HandshakeResolve(req.URL.Host)
	if nHost != "" {
		req.URL.Host = nHost
	}
	return req, nil
}

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Start Handshake Domain Proxy",
	Run: func(cmd *cobra.Command, args []string) {
		cache.Info()
		log.Println("Start Handshake proxy...")
		proxy := goproxy.NewProxyHttpServer()
		proxy.Verbose = true
		proxy.OnRequest().DoFunc(resoveHostDomain)
		log.Println("This Will start proxy :8080")
		log.Fatal(http.ListenAndServe(":8080", proxy))
	},
}
