package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"github.com/logologics/kunren-be/internal/route"
	"github.com/logologics/kunren-be/internal/api"
	d "github.com/logologics/kunren-be/internal/domain"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	// log config
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
	})

	viper.SetConfigName("application")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	address := viper.GetString("address")

	env := api.Env{}

	mux := route.NewRestRouter(&env)
	graceFul(&env)

	log.WithFields(log.Fields{
		"version": d.Version,
		"address": address,
		"repo":    viper.GetString("repo.type"),
	}).Info("Kunren started")

	https := &d.Https{}
	viper.UnmarshalKey("https", https)

	if https.Enabled {
		cert := https.CertPath
		key := https.KeyPath

		log.Fatal(http.ListenAndServeTLS(address, cert, key, mux))
	} else {
		log.Fatal(http.ListenAndServe(address, mux))
	}
}

func graceFul(env *api.Env) {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		log.Printf("caught sig: %+v", sig)
		os.Exit(0)
	}()
}
