package cmd

import (
	"fmt"
	"handshakeproxy/cache"
	"handshakeproxy/nextdns"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/json"
)

var (
	nextID     string
	proxyHost  string
	nextDNSTTL int32
	proxyPort  int32
	authFile   string
	noAuth     = false
	proxyCmd   = &cobra.Command{
		Use:   "proxy",
		Short: "Start Handshake Domain Proxy",
		Run:   proxyCmdRun,
	}
	userMaps = map[string]string{}
)

func setProxyFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&nextID, "nextid", "", "c3484e", "Define nextdns id")
	cmd.Flags().Int32VarP(&nextDNSTTL, "cache", "c", 300, "define nextdns dns cache seconds.")
	cmd.Flags().Int32VarP(&proxyPort, "port", "p", 8080, "proxy port listen.")
	cmd.Flags().StringVarP(&proxyHost, "host", "", "", "proxy bind host.")
	cmd.Flags().StringVarP(&authFile, "auth", "", "./passwd.json", "auth file path.")
}

func init() {
	setProxyFlags(proxyCmd)
	var (
		authBytes      []byte
		err            error
		authFileHandle *os.File
	)
	if authFileHandle, err = os.Open(authFile); err != nil {
		log.Println("Open auth file failed .  ", err)
		noAuth = true
	}

	if authBytes, err = ioutil.ReadAll(authFileHandle); err != nil {
		log.Println("auth file does not exist : ", authFile)
		noAuth = true
	}
	json.Unmarshal(authBytes, &userMaps)
}

func resoveHostDomain(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	log.Println(req.Header)
	nHost := nextdns.HandshakeResolve(req.URL.Host)
	if nHost != "" {
		req.URL.Host = nHost
	}
	return req, nil
}

func proxyCmdRun(cmd *cobra.Command, args []string) {
	cache.Init()
	nextdns.SetOptions(nextID, nextDNSTTL)
	log.Println("Start Handshake proxy... with noAuth: ", noAuth)
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	if !noAuth {
		auth.ProxyBasic(proxy, "my_realm", func(user, pwd string) bool {
			if p, exist := userMaps[user]; exist {
				return p == pwd
			}
			return false
		})
	}
	proxy.OnRequest().DoFunc(resoveHostDomain)
	bindStr := fmt.Sprintf("%s:%d", proxyHost, proxyPort)
	log.Println("This Will start proxy ", bindStr)

	log.Fatal(http.ListenAndServe(bindStr, proxy))
}
