package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {
	target := "https://docs.melonly.xyz"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := target + r.URL.Path
		if r.URL.RawQuery != "" {
			url += "?" + r.URL.RawQuery
		}
		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer resp.Body.Close()
		for k, v := range resp.Header {
			if k == "Content-Security-Policy" {
				newV := make([]string, len(v))
				for i, val := range v {
					newV[i] = strings.ReplaceAll(val, "frame-ancestors 'none';", "frame-ancestors *;")
				}
				w.Header()[k] = newV
			} else if k != "X-Frame-Options" {
				w.Header()[k] = v
			}
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(resp.StatusCode)
		if strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
			body, _ := io.ReadAll(resp.Body)
			bodyStr := string(body)
			bodyStr = strings.ReplaceAll(bodyStr, `<meta http-equiv="Content-Security-Policy"`, "")
			io.WriteString(w, bodyStr)
		} else {
			io.Copy(w, resp.Body)
		}
	})
	http.ListenAndServe(":8080", nil)
}
