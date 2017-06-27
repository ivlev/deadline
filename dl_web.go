package main

import (
"fmt"
"log"
"net/http"
)
func main() {
	http.HandleFunc("/", postHandler)
	log.Println("Deadline server start. Look yours dedlines by 127.0.0.1:3000")
	http.ListenAndServe(":3000", nil)
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Пока у вас всё в порядке")
}
