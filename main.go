package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("get-ip")

func getCurrentIP(r http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	if ip == "[" {
		ip = "127.0.0.1"
	}
	return ip
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := getCurrentIP(*r)
		log.Info(ip)
		io.WriteString(w, ip)
	})
	log.Info("Start Server in 12366...")
	http.ListenAndServe(":12366", nil)
}
