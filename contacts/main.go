package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	PORT string
)

func init() {
	PORT = os.Getenv("PORT")
}

func init() {
	if PORT == "" {
		flag.StringVar(&PORT, "port", "50080", "--port=50080")
		flag.Parse()
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})
	http.HandleFunc("/ping", pong)
	http.HandleFunc("/greet", greet("Hello Folks!"))
	health := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ok")
	}
	http.HandleFunc("/health", health)
	fmt.Println("Server started on port " + PORT)
	http.ListenAndServe(":"+PORT, nil) // on all interfaces.
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pong")
}
func greet(msg string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, msg)
	}
}

// :50050 what is the ip address?
// 0.0.0.0 != 127.0.0.1
// 127.0.0.1
// 192.168.0.4
// 10.45.0.4
// 192.168.0.1/24 -- 192.168.0.1 --- 192.168.0.255
// netwrok interfaces
// LO,ETH0,WSLP,
// If your computer is connectec through a model to a wireless router --> There must be an interface
// loopback--> connecting to your own system.
// eithernet interface
