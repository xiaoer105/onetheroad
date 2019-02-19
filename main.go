package main

import (
	"net/http"
	"app/org/web/OneTheRoad/routers"
	"time"
	"log"
)

func main() {

	server := &http.Server{
		Addr:         ":8080",        
		Handler:      routers.AppTLSRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}

}
