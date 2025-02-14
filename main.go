package main

import (
	"time"

	"TCP_Server_1/myServer"
)

func main() {
	srv := myServer.Server{
		Addr:         ":8080",
		IdleTimeout:  10 * time.Second,
		MaxReadBytes: 1000,
	}
	go srv.ListenAndServe()
	time.Sleep(10 * time.Second)
	//srv.Shutdown()
	for {
	}
}
