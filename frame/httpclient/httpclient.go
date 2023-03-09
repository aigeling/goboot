package httpclient

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"time"
)

type HttpClientConfig struct {
	ProxyUrl string
	Timeout  int
}

// 客户端初始化
func NewHttpClient(conf *HttpClientConfig) *http.Client {

	var proxy func(*http.Request) (*url.URL, error) = nil

	if conf.ProxyUrl != "" {
		proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse(conf.ProxyUrl)
		}
	}
	transport := &http.Transport{
		Proxy: proxy,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 60 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          20,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 10 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}

	if conf.Timeout == 0 {
		conf.Timeout = 60
	}

	return &http.Client{
		Transport: transport,
		Timeout:   time.Duration(conf.Timeout) * time.Second,
	}
}
