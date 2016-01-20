package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var status string
	_, err := http.Get("https://web.nvd.nist.gov/view/800-53/home")
	if err != nil {
		status = "Failed"
	}
	status = "Success"
	fmt.Fprintf(w, "%s", status)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
