package main

import (
	"NATS_reader/pkg/events"
	"log"
	"net"
	"net/http"

	"NATS_reader/cmd/app"
	"NATS_reader/pkg/configs"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var config configs.Config

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalf("configs err: %s", err.Error())
	}
	log.Println("configs -> Done!")
	router := mux.NewRouter()
	initStan := events.NewEvent(config)
	stan, err := initStan.InitStan(config)
	if err != nil {
		log.Fatalf("initStan.Error: %v", err)
	}
	server := app.NewServer(router, config.Prefix, stan)
	server.InitRoutes()
	addr := net.JoinHostPort(config.Host, config.Port)
	log.Printf("server will start at address: %s", addr)
	log.Printf("endPOINT address: %s", addr+"/"+config.Prefix)

	handler := cors.Default().Handler(server)
	loggedRouter := handlers.LoggingHandler(log.Writer(), handler)

	if err = http.ListenAndServe(addr, loggedRouter); err != nil {
		log.Println(err)
		return
	}
}
