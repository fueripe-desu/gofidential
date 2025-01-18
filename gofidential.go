package gofidential

import (
	loader "github.com/fueripe-desu/gofidential/internal/loader"
	parser "github.com/fueripe-desu/gofidential/internal/parser"
	reflector "github.com/fueripe-desu/gofidential/internal/reflector"
)

type Environment struct {
	Name           string
	OverridePath   string
	IgnoreFilename bool
}

func Load(s any, env Environment) error {
	byteData, err := loader.Load(env.Name, env.OverridePath, env.IgnoreFilename)

	if err != nil {
		return err
	}

	parsedData, err := parser.Parse(byteData)

	if err != nil {
		return err
	}

	return reflector.Reflect(parsedData, s)
}
