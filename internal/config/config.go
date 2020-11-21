package config

import (
	"os"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/logologics/koanf/providers/s3"
	"github.com/logologics/kunren-be/internal/domain"
	"github.com/sirupsen/logrus"
)

// Load loads the config from file system
func Load() (*domain.Config, error) {
	k := koanf.New(".")
	conf := &domain.Config{}

	if err := k.Load(file.Provider("./config/application.json"), json.Parser()); err != nil {
		logrus.Errorf("error loading config: ./config/application.json %v", err)
	} else {
		logrus.Infof("loading config: from ./config/application.json %v", err)
		k.Unmarshal("", &conf)
		return conf, nil
	}

	if err := k.Load(file.Provider("/config/application.json"), json.Parser()); err != nil {
		logrus.Errorf("error loading config: /config/application.json %v", err)
	} else {
		logrus.Infof("loading config from  /config/application.json %v", err)
		k.Unmarshal("", &conf)
		return conf, nil
	}

	p, err := s3.Provider(s3.Config{
		Region:    os.Getenv("AWS_S3_REGION"),
		Bucket:    os.Getenv("AWS_S3_BUCKET"),
		ObjectKey: "kunren/config/config.json",
		UseIAM:    true,
	})
	if err != nil {
		logrus.Errorf("error connecting to bucket %v in %v: %v", os.Getenv("AWS_S3_BUCKET"), os.Getenv("AWS_S3_REGION"), err)
		return conf, err
	}

	if err := k.Load(p, json.Parser()); err != nil {
		logrus.Errorf("error loading config from s3 bucket %v: %v", os.Getenv("AWS_S3_BUCKET"), err)
	} else {
		logrus.Errorf("error loading config from s3 bucket %v: %v", os.Getenv("AWS_S3_BUCKET"), err)
		k.Unmarshal("", &conf)
		return conf, nil
	}

	return conf, nil
}
