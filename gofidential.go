// Package gofidential provides a simple, flexible, and scalable interface
// for loading .env files and managing different environments.
//
// It exports the [gofidential.Load] function, which loads a .env file into
// a struct, and the [gofidential.Environment] struct, which allows users
// to configure different environments.
//
// This package is ideal for applications of all sizes, from small hobby
// projects to full-blown enterprise apps, especially applications that
// require a robust secrets loader that supports environment-specific
// configurations (such as development, testing, and production), and can
// also scale seamlessly. It also includes built-in type validation to
// ensure the environment variables meet the required value criteria.
//
// The library itself enforces strict parsing rules which makes the .env
// files much more predictable and free of bugs. The parsing rules were
// defined by the standard created in the library itself, the GoFidential
// Standard v1 (GFSv1), and it is supposed to define a new and more predictable
// set of rules for creating .env files with all the best practices.
//
// Example usage:
//
//	package main
//
//	import (
//		"os"
//
//		gf "github.com/fueripe-desu/gofidential"
//	)
//
//	type Secrets struct {
//		ApiKey 		string
//		BaseUrl 	string
//		LogLevel 	int
//	}
//
//	int main() {
//		env := gf.Environment{
//			Name: "dev",
//			IgnoreFilename: true,
//		}
//
//		s := Secrets{}
//		err := gf.Load(&s, env)
//
//		if err != nil {
//			castErr := err.(*gf.GofidentialError)
//			log.Println(castErr.Message())
//			log.Println(castErr.Hint())
//			os.Exit(-1)
//		}
//
//		log.Printf("ApiKey: '%s'\n", s.ApiKey)
//		log.Printf("BaseUrl: '%s'\n", s.BaseUrl)
//		log.Printf("LogLevel: '%d'\n", s.LogLevel)
//	}
package gofidential

import (
	loader "github.com/fueripe-desu/gofidential/internal/loader"
	parser "github.com/fueripe-desu/gofidential/internal/parser"
	reflector "github.com/fueripe-desu/gofidential/internal/reflector"
)

// The [gofidential.Environment] struct is used to configure an environment.
//
// Fields:
//   - Name (string): The name of the environment. This field is required.
//   - OverridePath (string): An optional custom directory path to look for the .env file.
//   - IgnoreFilename (bool): Determines the naming convention for the .env file.
//     If true, the filename will always be ".env". If false, the filename will include
//     the environment name (e.g., "dev.env", "test.env", "prod.env").
//
// Constraints:
//   - "Name" must not be empty.
//   - "Name" must only contain lowercase letters and underscores.
//   - "Name" must not have leading or trailing underscores.
//   - "Name" must not consist solely of underscores.
//   - If specified, "OverridePath" must be a valid directory path (not a filepath).
//
// Notes:
//   - Both "OverridePath" and "IgnoreFilename" are optional.
type Environment struct {
	Name           string
	OverridePath   string
	IgnoreFilename bool
}

// The [gofidential.Load] function loads .env file data into a specified struct.
//
// Parameters:
//   - s (any): A pointer to a struct that will be populated with the .env file data.
//   - env (Environment): The environment configuration.
//
// Returns:
//   - error: An error describing why the application failed to load the .env file.
//
// Constraints:
//   - "s" must not be nil.
//   - "s" must be a pointer.
//   - "s" must point to a struct (not a primitive type, slice, map, etc.).
//   - "env" must not be nil.
//
// Usage example:
//
//	env := gf.Environment{
//		Name:           "dev",
//		IgnoreFilename: true,
//	}
//
//	s := Secrets{}
//	err := gf.Load(&s, env)
//	if err != nil {
//		// Cast the error to *gf.GofidentialError for detailed information
//		castErr := err.(*gf.GofidentialError)
//		log.Println(castErr.Message())
//		log.Println(castErr.Hint())
//		os.Exit(-1)
//	}
//
//	log.Printf("ApiKey: '%s'\n", s.ApiKey)
//	log.Printf("BaseUrl: '%s'\n", s.BaseUrl)
//	log.Printf("LogLevel: '%d'\n", s.LogLevel)
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
