package cmd

import (
	"fmt"
	"handshakeproxy/cache"
	"handshakeproxy/nextdns"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/spf13/cobra"
)

var (
	nextID     string
	proxyHost  string
	nextDNSTTL int32
	proxyPort  int32
	proxyCmd   = &cobra.Command{
		Use:   "proxy",
		Short: "Start Handshake Domain Proxy",
		Run:   proxyCmdRun,
	}
)

func setProxyFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&nextID, "nextid", "", "c3484e", "Define nextdns id")
	cmd.Flags().Int32VarP(&nextDNSTTL, "cache", "c", 300, "define nextdns dns cache seconds.")
	cmd.Flags().Int32VarP(&proxyPort, "port", "p", 8080, "proxy port listen.")
	cmd.Flags().StringVarP(&proxyHost, "host", "", "", "proxy bind host.")
}

func init() {
	setProxyFlags(proxyCmd)
}

func resoveHostDomain(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	nHost := nextdns.HandshakeResolve(req.URL.Host)
	if nHost != "" {
		req.URL.Host = nHost
	}
	return req, nil
}

func proxyCmdRun(cmd *cobra.Command, args []string) {
	cache.Init()
	nextdns.SetOptions(nextID, nextDNSTTL)
	log.Println("Start Handshake proxy...")
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(resoveHostDomain)
	bindStr := fmt.Sprintf("%s:%d", proxyHost, proxyPort)
	log.Println("This Will start proxy ", bindStr)

	log.Fatal(http.ListenAndServe(bindStr, proxy))
}
