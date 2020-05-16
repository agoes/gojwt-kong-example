package main

import (
	"log"
	"net/http"

	route "github.com/agoes/jwt-issuer/route"
	"github.com/go-chi/chi"
)

func main() {
	router := route.Routes()

	log.Println("Initialize routes ...")
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Route error : %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":8080", router))
	// fmt.Println(token.CreateToken("key"))
	// fmt.Println(kong.GetConsumerJwtCredentials("john.doe@carisini.local", "key"))
}
