package main

import (
	"app/org/web/OnTheRoad/routers"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         ":80",
		Handler:      routers.AppTLSRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}

}
