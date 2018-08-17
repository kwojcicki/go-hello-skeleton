package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/zenazn/goji/graceful"
	"go-hello-skeleton/rest"
	"syscall"
	"time"
)

func main() {
	r := rest.InitRouter()
	log := logrus.New()

	graceful.AddSignal(syscall.SIGINT, syscall.SIGTERM)
	graceful.Timeout(30 * time.Second)
	graceful.PreHook(func() {
		log.Warnln("Shutdown initiated.")
	})
	graceful.PostHook(func() {
		log.Warnln("Shutdown completed.")
	})

	server := &graceful.Server{
		Addr:    "0.0.0.0:8080",
		Handler: chi.ServerBaseContext(context.Background(), r),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}

	graceful.Wait()
}
