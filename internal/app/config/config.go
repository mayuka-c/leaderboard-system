package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

var (
	serviceConfig = ServiceConfig{}
	dbConfig      = DBConfig{}
)

type ServiceConfig struct {
	APIPort int `envconfig:"PORT" default:"8181"`
}

type DBConfig struct {
	DB_URL   string `envconfig:"DB_URL" default:"localhost:5432"`
	Username string `envconfig:"DB_USER" default:"root"`
	Password string `envconfig:"DB_USER" default:"mysecretpassword"`
	Database string `envconfig:"DB_Database" default:"leaderboard-system"`
}

func LoadServiceConfig() {
	log.Println("Load Service configuration")
	err := envconfig.Process("e-commerce", &serviceConfig)
	if err != nil {
		log.Fatalln("Failed fetching service configs")
		panic(err)
	}
}

func LoadDBConfig() {
	log.Println("Load DB configuration")
	err := envconfig.Process("e-commerce", &dbConfig)
	if err != nil {
		log.Fatalln("Failed fetching db configs")
		panic(err)
	}
}

// GetServiceConfig method to fetch the ServiceConfig
func GetServiceConfig() ServiceConfig {
	return serviceConfig
}

// GetDBConfig get db env vars
func GetDBConfig() DBConfig {
	return dbConfig
}
