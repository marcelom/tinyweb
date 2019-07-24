package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

var (
	dirPtr  = flag.String("dir", ".", "the directory with files to be served")
	portPtr = flag.Int("port", 8080, "the port to listen to")
)

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dirPtr)))
	http.HandleFunc("/healthz", healthcheck)

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Listening on the following IP addresses:")
	for _, addr := range addrs {
		a := strings.Split(addr.String(), "/") // split the addr/mask pair
		fmt.Printf(" * http://%v:%d\n", a[0], *portPtr)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), nil); err != nil {
		log.Fatal(err)
	}
}
