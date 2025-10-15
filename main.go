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
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		req.Header.Set("Accept-Language", "en-US,en;q=0.5")
		req.Header.Set("Accept-Encoding", "gzip, deflate")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
	}
	proxy.ModifyResponse = func(resp *http.Response) error {
		resp.Header.Del("X-Frame-Options")
		resp.Header.Del("Content-Security-Policy")
		resp.Header.Set("Access-Control-Allow-Origin", "*")
		return nil
	}
	http.Handle("/", proxy)
	http.ListenAndServe(":8080", nil)
}
