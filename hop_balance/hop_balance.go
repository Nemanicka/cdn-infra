package main

import (
    "net/http"
    "net/http/httputil"
    "net/url"
    "log"
    "strings"
    "fmt"
    "os/exec"
)

var server1 = "http://127.0.0.1:5002"
var server2 = "http://127.0.0.1:5003"

func getBestHopServer() string {
    h1, _ := countHops(server1)
    h2, _ := countHops(server2)
    fmt.Println(h1, h2)

    if h1 > h2 {
        return server2
    }
    return server1
}

func countHops(target string) (int, error) {
	cmd := exec.Command("traceroute", target)
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(output), "\n")
	return len(lines) - 1, nil
} 

func main() {
    targetUrl, _ := url.Parse( getBestHopServer() )

    proxy := httputil.NewSingleHostReverseProxy(targetUrl)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        proxy.ServeHTTP(w, r)
    })

    log.Fatal(http.ListenAndServe("127.0.0.1:82", nil))
}
