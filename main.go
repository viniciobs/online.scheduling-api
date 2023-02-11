package main

import (
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/online.scheduling-api/router"
	"github.com/sarulabs/di"

	_di "github.com/online.scheduling-api/di"
)

func main() {
	log.Println("Initializing...")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	builder, _ := di.NewBuilder()
	builder.Add(_di.Services...)

	app := builder.Build()
	defer app.Delete()

	router := router.ConfigureRouter(app)

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
