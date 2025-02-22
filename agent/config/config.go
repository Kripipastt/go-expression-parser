package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	Port            string
	ComputingPower  int
	OrchestratorUrl string
}

func LoadConfig() *Config {
	godotenv.Load()
	conf := new(Config)

	conf.Port = os.Getenv("PORT")
	if conf.Port == "" {
		conf.Port = "8081"
	}

	computingPower := os.Getenv("COMPUTING_POWER")
	if computingPower != "" {
		conf.ComputingPower, _ = strconv.Atoi(computingPower)
	} else {
		conf.ComputingPower = 3
	}

	conf.OrchestratorUrl = os.Getenv("ORCHESTRATOR_URL")
	if conf.OrchestratorUrl == "" {
		panic("ORCHESTRATOR_URL environment variable not set")
	}

	return conf
}

var Service = LoadConfig()
