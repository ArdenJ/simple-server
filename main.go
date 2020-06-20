package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ardenj/graceful-http/handlers"
)

var Port = "5678"

func main() {
	l := log.New(os.Stdout, "api", log.LstdFlags)

	h := handlers.NewHowdy(l)
	b := handlers.NewBoop(l)
	p := handlers.NewProducts(l)

	smux := http.NewServeMux()
	smux.Handle("/", h)
	smux.Handle("/beep", b)
	smux.Handle("/products", p)

	s := &http.Server{
		Addr:         ":" + Port, // set TCP address to listen on
		Handler:      smux,       // handler to invoke
		ErrorLog:     l,          // logger for errors while accepting connections
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Println("Server started on port: 5678")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	defer s.Shutdown(tc)

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)

	// sig blocks until it receives a message from the channel.
	// This prevents the server from shutting down immediately
	sig := <-signalChan
	l.Println("Signal received: gracefully shutting this thing down \n%s", sig)
}
