package config

import (
	"fmt"

	"github.com/logologics/kunren-be/internal/domain"
	"github.com/spf13/viper"
)

// Load loads the config from file system
func Load() (*domain.Config, error) {
	viper.SetConfigName("application")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() 
	if err != nil {            
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	conf := &domain.Config{}
	if err := viper.Unmarshal(conf); err != nil {
		return &domain.Config{}, err
	}

	return conf, nil
}
