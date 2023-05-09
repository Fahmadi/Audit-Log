package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	GrpcURI            string `mapstructure:"GRPC_URI"`
	HttpURI            string `mapstructure:"HTTP_URI"`
	DatabaseDriverName string `mapstructure:"DATABASE_DRIVER_NAME"`
	DatabaseURI        string `mapstructure:"DATABASE_URI"`
	BrokerURI          string `mapstructure:"BROKER_URI"`
	AutoMigrate        bool   `mapstructure:"AUTO_MIGRATE"`
}

func LoadConfig() config {
	v := viper.New()

	v.AddConfigPath("./")
	v.AddConfigPath("./configs/")
	v.SetConfigName("auditLog")
	v.SetConfigType("env")
	v.SetEnvPrefix("AUDITLOG")
	v.AutomaticEnv()
	v.BindEnv("GRPC_URI")
	v.BindEnv("HTTP_URI")
	v.BindEnv("DATABASE_URI")
	v.BindEnv("BROKER_URI")

	if err := v.ReadInConfig(); err != nil {
		log.Println("Couldn't find the .env config file")
	}

	c := config{}
	if err := v.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}

	log.Printf("config: %+v", c)

	return c
}
