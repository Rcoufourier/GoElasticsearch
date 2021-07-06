package controllers

import (
	"fmt"
	"log"
	"net/http"
)


func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "welcome on the search page blablablu")
	log.Println("SearchBooks")
}
