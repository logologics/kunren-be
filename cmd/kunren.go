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
		DisableColors: false,
	})

	viper.SetConfigName("application")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	config := &d.Config{}
	viper.Unmarshal(config)
	repo, err := api.CreateRepo(config)
	if err != nil {
		log.Fatal("Can't create repo: %v", err)
	}
	env := api.Env{Config: config, Repo: repo}

	mux := route.NewRestRouter(&env)
	graceFul(&env)

	log.WithFields(log.Fields{
		"version": d.Version,
		"address": config.Address,
	}).Info("Kunren started")

	if config.Https.Enabled {
		cert := config.Https.CertPath
		key := config.Https.KeyPath

		log.Fatal(http.ListenAndServeTLS(config.Address, cert, key, mux))
	} else {
		log.Fatal(http.ListenAndServe(config.Address, mux))
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
