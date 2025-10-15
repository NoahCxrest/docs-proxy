package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, _ := url.Parse("https://docs.melonly.xyz")
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Director = func(req *http.Request) {
		req.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
	}
	proxy.ModifyResponse = func(resp *http.Response) error {
		return nil
	}
	http.Handle("/", proxy)
	http.ListenAndServe(":8080", nil)
}
