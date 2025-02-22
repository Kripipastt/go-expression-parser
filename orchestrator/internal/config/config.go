package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	Port                 string
	TimeAdditionMs       int
	TimeSubtractionMs    int
	TimeMultiplicationMs int
	TimeDivisionMs       int
	TimeExponentiationMs int
}

func LoadConfig() *Config {
	godotenv.Load()
	conf := new(Config)
	conf.Port = os.Getenv("PORT")
	if conf.Port == "" {
		conf.Port = "8080"
	}

	timeAdditionMs := os.Getenv("TIME_ADDITION_MS")
	if timeAdditionMs != "" {
		conf.TimeAdditionMs, _ = strconv.Atoi(timeAdditionMs)
	} else {
		conf.TimeAdditionMs = 1000
	}
	timeSubtractionMs := os.Getenv("TIME_SUBTRACTION_MS")
	if timeSubtractionMs != "" {
		conf.TimeSubtractionMs, _ = strconv.Atoi(timeSubtractionMs)
	} else {
		conf.TimeSubtractionMs = 1000
	}
	timeMultiplicationMs := os.Getenv("TIME_MULTIPLICATION_MS")
	if timeMultiplicationMs != "" {
		conf.TimeMultiplicationMs, _ = strconv.Atoi(timeMultiplicationMs)
	} else {
		conf.TimeMultiplicationMs = 3000
	}
	timeDivisionMs := os.Getenv("TIME_DIVISION_MS")
	if timeDivisionMs != "" {
		conf.TimeDivisionMs, _ = strconv.Atoi(timeDivisionMs)
	} else {
		conf.TimeDivisionMs = 3000
	}
	timeExponentiation := os.Getenv("TIME_EXPONENTIATION_MS")
	if timeExponentiation != "" {
		conf.TimeExponentiationMs, _ = strconv.Atoi(timeExponentiation)
	} else {
		conf.TimeExponentiationMs = 5000
	}

	return conf
}

var Service = LoadConfig()
