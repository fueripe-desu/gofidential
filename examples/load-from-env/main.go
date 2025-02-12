package main

import (
	"fmt"
	"log"
	"os"

	gf "github.com/fueripe-desu/gofidential/v2"
)

type Secrets struct {
	AppName string
	Version int
}

func main() {
	os.Setenv("APP_NAME", "Gofidential")
	os.Setenv("VERSION", "1")

	env := gf.Environment{Name: "dev", LoadFromEnv: true}
	s := Secrets{}

	err := gf.Load(&s, env)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("App name: '%v'\n", s.AppName)
	fmt.Printf("Version: '%v'\n", s.Version)

	os.Unsetenv("APP_NAME")
	os.Unsetenv("VERSION")
}
