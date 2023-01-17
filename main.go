package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/online.scheduling-api/src/ioc"
	"github.com/online.scheduling-api/src/router"
	"github.com/sarulabs/di"
)

func main() {
	builder, _ := di.NewBuilder()
	builder.Add(ioc.Services...)

	app := builder.Build()
	defer app.Delete()

	router := router.ConfigureRouter(app)

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  3 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err.Error())
	}
}
