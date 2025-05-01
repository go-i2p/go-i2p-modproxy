package main

import (
	"net"
	"net/http"
	"time"

	"github.com/go-i2p/onramp"
	"github.com/goproxy/goproxy"

	metadialer "github.com/go-i2p/go-meta-dialer"
)

func Dial(network, addr string) (net.Conn, error) {
	return metadialer.Dial(network, addr)
}

func init() {
	// Initialize the client with a default timeout
	http.DefaultClient = &http.Client{
		Timeout: 360 * time.Second,
		Transport: &http.Transport{
			Dial: Dial,
		},
	}
}

func main() {
	if garlic, err := onramp.NewGarlic("go-i2p-modules-proxy", "127.0.0.1:7656", onramp.OPT_DEFAULTS); err != nil {
		panic(err)
	} else {
		fetcher := &goproxy.GoFetcher{
			Transport: &http.Transport{
				Dial: Dial,
			},
		}
		listener, err := garlic.Listen()
		if err != nil {
			panic(err)
		}
		defer listener.Close()
		http.Serve(listener, &goproxy.Goproxy{
			Fetcher: fetcher,
		})
	}
}
