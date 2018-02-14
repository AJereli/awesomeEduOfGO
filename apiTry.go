package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"log"
	"net/http"

)




func main() {
	fmt.Println("Server start\nPress Ctrl+C to shutdown")
	router := InitRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}


	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)


	<-c

	var wait time.Duration
	wait = time.Second * 10
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)

}

