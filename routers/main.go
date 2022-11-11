package routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHomeRoutes(router *mux.Router) {
	router.HandleFunc("/", welcomePage).Methods("GET")
}

func welcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This website is used for educational purposes only!")
}
