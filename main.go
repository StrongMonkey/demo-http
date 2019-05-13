package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)



func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	v := int(math.Mod(float64(rand.Intn(100)), 4))
	switch v {
	case 0:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "200 code")
	case 1:
		w.WriteHeader(http.StatusMovedPermanently)
		fmt.Fprintln(w, "301 code")
	case 2:
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "401 code")
	case 3:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "500 code")
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
