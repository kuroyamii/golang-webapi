package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	Address string
	Handler http.Handler
}

//Listen and serve method
func (s Server) ListenAndServe() {
	go func() {
		err := http.ListenAndServe(s.Address, s.Handler)
		if err != nil {
			log.Printf("ERROR (starting server): %v\n", err)
		}
	}()
	log.Printf("INFO (server): server started, listening to %s\n", s.Address)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("INFO (server): shutting down the server")
}

//Provide the server
func ProvideServer(address string, handler http.Handler) Server {
	return Server{
		Address: address,
		Handler: handler,
	}
}
