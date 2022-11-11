package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/followthepattern/go-simple-webservice/routers"
	"github.com/followthepattern/go-simple-webservice/services"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	port := "8080"
	fmt.Println("FP simple webservice is running!")
	fmt.Println("Listening on " + port)

	router := mux.NewRouter()
	database := services.NewDatabase()

	routers.RegisterHomeRoutes(router)
	routers.NewUserRouterRegister(database).RegisterUserRoutes(router)

	originsAllowed := handlers.AllowedOrigins([]string{"*"})
	methodsAllowed := handlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions})
	headersAllowed := handlers.AllowedHeaders([]string{"application/json", "Content-Type", "x-correlationid", "X-Requested-With", "authorization"})

	log.Fatal(http.ListenAndServe(":"+port,
		handlers.CORS(originsAllowed, headersAllowed, methodsAllowed)(router)))
}
