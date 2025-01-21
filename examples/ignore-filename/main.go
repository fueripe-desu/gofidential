package main

import (
	"fmt"
	"log"

	gf "github.com/fueripe-desu/gofidential"
)

type Secrets struct {
	// Application configuration
	AppName     string
	Version     int
	Debug       bool
	Environment string
}

func main() {
	env := gf.Environment{
		Name: "dev",
		// This will make the application look for the ".env" filename instead of "dev.env".
		IgnoreFilename: true,
	}
	s := Secrets{}

	err := gf.Load(&s, env)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("App name: '%v'\n", s.AppName)
	fmt.Printf("Version: '%v'\n", s.Version)
	fmt.Printf("Debug: '%v'\n", s.Debug)
	fmt.Printf("Environment: '%v'\n", s.Environment)
}
