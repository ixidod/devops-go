package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "jak sie masz")
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server smiga na porcie 8080")
	http.ListenAndServe(":8080", nil)
}
