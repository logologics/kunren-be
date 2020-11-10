package main

import (
	"github.com/logologics/kunren-be/internal/config"
	d "github.com/logologics/kunren-be/internal/domain"
	"github.com/logologics/kunren-be/internal/route"

	"github.com/sirupsen/logrus"
)

func main() {

	// log config
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
	})

	conf, err := config.Load()
	if err != nil {
		logrus.Fatalf("Loading config failed %v", err)
	}
	router, err := route.New(conf)
	router.Run(":9876")

	logrus.Infof("Kunren %v started at %v ", d.Version, conf.Address)

}
