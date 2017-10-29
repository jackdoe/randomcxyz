package main

import (
	"net/http"
	"strconv"
	"math/rand"
	"fmt"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		n, err := strconv.Atoi(r.URL.Query().Get("n"))
		if (err != nil || n < 4) {
			n = 4
		}

		if (r.URL.Query().Get("octet") == "stream") {
			w.Header().Set("Content-Type", "octet/stream")
			w.Header().Set("Content-Disposition", "attachment; filename=randomcxyz.c")
		} else {
			w.Header().Set("Content-Type", "text/plain")
		}
		w.Write([]byte("unsigned int randomcxyz[] = {"))
		for i := 0; i < n; i+= 1 {
			if i % 4 == 0 {
				w.Write([]byte("\n\t"))
			}
			if (i == n - 1) {
				w.Write([]byte(fmt.Sprintf("0x%08x", rand.Int31())))
			} else {
				w.Write([]byte(fmt.Sprintf("0x%08x,", rand.Int31())))
			}
		}
		w.Write([]byte("\n};"))
	})
	http.ListenAndServe(":4569", nil)
}
