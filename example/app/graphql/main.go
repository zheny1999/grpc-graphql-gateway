package main

import (
	"log"
	"net/http"

	"github.com/ysugimoto/grpc-graphql-gateway/example/spec/starwars"
	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
)

func main() {
	mux := runtime.NewServeMux(runtime.Cors())

	if err := starwars.RegisterStartwarsServiceGraphqlHandler(mux, nil); err != nil {
		log.Fatalln(err)
	}

	http.Handle("/graphql", mux)
	log.Fatalln(http.ListenAndServe(":8888", nil))
}
