package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/golang/glog"
)

func handler(res http.ResponseWriter, req *http.Request) {
	flag.Set("v", "5")
	for k, v := range req.Header {
		res.Header().Set(k, v[0])
	}
	res.Header().Set("version", runtime.Version()+"/"+runtime.GOOS+"/"+runtime.GOARCH)
	res.WriteHeader(http.StatusOK)
	glog.V(5).Info("IP:", req.RemoteAddr, "\t", "httpStatus:", "200")
}
func health(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, 200)
	res.WriteHeader(http.StatusOK)
}

func server() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", health)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	server()
}
