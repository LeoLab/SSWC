package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat("/store/static/"); err == nil {
	} else {
	}
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {

	if _, err := os.Stat("/store/res/"); err == nil {
	} else {
	}
}
*/

func main() {
	/*
		http.HandleFunc("/", defaultHandler)
		http.HandleFunc("/res/", resourceHandler)
	*/
	fs := http.FileServer(http.Dir("/static"))
	http.Handle("/", fs)
	fmt.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
