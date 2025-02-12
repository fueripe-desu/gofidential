package main

import (
	"fmt"
	"log"

	gf "github.com/fueripe-desu/gofidential/v2"
)

type Secrets struct {
	// Application configuration
	AppName     string
	Version     int
	Debug       bool
	Environment string

	// Database configuration
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string

	// Logging
	LogFile  string
	LogLevel string

	// API configuration
	ApiUrl string
	ApiKey string
}

func main() {
	env := gf.Environment{Name: "dev"}
	s := Secrets{}

	err := gf.Load(&s, env)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("App name: '%v'\n", s.AppName)
	fmt.Printf("Version: '%v'\n", s.Version)
	fmt.Printf("Debug: '%v'\n", s.Debug)
	fmt.Printf("Environment: '%v'\n", s.Environment)
	fmt.Printf("Database host: '%v'\n", s.DatabaseHost)
	fmt.Printf("Database port: '%v'\n", s.DatabasePort)
	fmt.Printf("Database name: '%v'\n", s.DatabaseName)
	fmt.Printf("Database user: '%v'\n", s.DatabaseUser)
	fmt.Printf("Database password: '%v'\n", s.DatabasePassword)
	fmt.Printf("Log file: '%v'\n", s.LogFile)
	fmt.Printf("Log level: '%v'\n", s.LogLevel)
	fmt.Printf("Api url: '%v'\n", s.ApiUrl)
	fmt.Printf("Api key: '%v'\n", s.ApiKey)
}
