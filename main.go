package main

import (
	"log"
	"net/http"
	"time"

	"github.com/online.scheduling-api/ioc"
	"github.com/sarulabs/di"
)

func main() {
	log.Println("Initializing...")

	builder, _ := di.NewBuilder()
	builder.Add(ioc.Services...)

	app := builder.Build()
	defer app.Delete()

	router := ConfigureRouter(app)

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  3 * time.Second,
	}

	log.Printf("Serving on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Print("Shutting down...")
		log.Fatalln(err.Error())
	}
}
