package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello Wrold")
	})
	http.ListenAndServe(":50089", nil)
}
