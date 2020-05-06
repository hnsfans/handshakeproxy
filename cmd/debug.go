package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/cobra"
)

var (
	debugCmd = &cobra.Command{
		Use:   "debug",
		Short: "Send Debug Request. Just Request http://welcomd.2d from proxy",
		Run:   debugRun,
	}
)

func init() {
	debugCmd.Flags().Int32VarP(&proxyPort, "port", "p", 8080, "proxy port listen.")
	debugCmd.Flags().StringVarP(&proxyHost, "host", "", "", "proxy bind host.")
}
func debugRun(cmd *cobra.Command, args []string) {
	var (
		err      error
		resp     *http.Response
		proxyURL *url.URL
	)

	if proxyHost == "" {
		proxyHost = "127.0.0.1"
	}
	proxy := fmt.Sprintf("http://%s:%d", proxyHost, proxyPort)
	log.Println("Debug Request ... Proxy to -> ", proxy)
	if proxyURL, err = url.Parse(proxy); err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
		Timeout: time.Second * 5, //超时时间
	}

	if resp, err = client.Get("http://welcome.2d"); err == nil {
		defer resp.Body.Close()
		if bodyBytes, err := ioutil.ReadAll(resp.Body); err == nil {
			log.Println(string(bodyBytes))
		}
	} else {
		log.Println(err)
	}
}
