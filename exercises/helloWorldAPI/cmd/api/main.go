package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// GetRoot Handler for Hello World return message
func GetRoot(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	w.WriteHeader(status)
	_, err := w.Write([]byte(`{"message": "Hello World"}`))

	if err != nil {
		log.WithFields(log.Fields{
			"method": r.Method,
			"route":  r.RequestURI,
			"error":  err,
		}).Warn()
	} else {
		log.WithFields(log.Fields{
			"method": r.Method,
			"route":  r.RequestURI,
			"status": status,
		}).Info()
	}
}

// NotFound Handler for endpoints that do not exist
func NotFound(w http.ResponseWriter, r *http.Request) {
	status := http.StatusNotFound
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte(`{"message": "not found"}`))

	if err != nil {
		log.WithFields(log.Fields{
			"method": r.Method,
			"route":  r.RequestURI,
			"error":  err,
		}).Warn()
	} else {
		log.WithFields(log.Fields{
			"method": r.Method,
			"route":  r.RequestURI,
			"status": status,
		}).Warn()
	}
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/hello", GetRoot).Methods(http.MethodGet)
	r.PathPrefix("").HandlerFunc(NotFound)

	addr := "0.0.0.0:8080"
	srv := &http.Server{
		Addr: addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Info("Listen on ", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err := srv.Shutdown(ctx); err != nil {
		log.Error(err)
	}
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info("Shutting down")
	os.Exit(0)
}
