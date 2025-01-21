package gofidential

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Load(t *testing.T) {
	t.Run("should load the env file correctly", func(t *testing.T) {
		require := require.New(t)
		assert := assert.New(t)

		s := struct {
			AppName     string
			Version     int
			Debug       bool
			Environment string

			DatabaseHost     string
			DatabasePort     string
			DatabaseName     string
			DatabaseUser     string
			DatabasePassword string

			LogFile  string
			LogLevel string

			ApiUrl string
			ApiKey string
		}{}

		expectedS := struct {
			AppName     string
			Version     int
			Debug       bool
			Environment string

			DatabaseHost     string
			DatabasePort     string
			DatabaseName     string
			DatabaseUser     string
			DatabasePassword string

			LogFile  string
			LogLevel string

			ApiUrl string
			ApiKey string
		}{
			AppName:     "MyApp",
			Version:     1,
			Debug:       true,
			Environment: "production",

			DatabaseHost:     "localhost",
			DatabasePort:     "5432",
			DatabaseName:     "my_database",
			DatabaseUser:     "db_user",
			DatabasePassword: "secure_password",

			LogFile:  "app_log.txt",
			LogLevel: "info",

			ApiUrl: "https://api.example.com",
			ApiKey: "your_api_key_here",
		}

		env := Environment{
			Name:         "example",
			OverridePath: "testdata/gofidential",
		}

		err := Load(&s, env)
		require.NoError(err, "An unexpected error occured")

		assert.Equal(s, expectedS)
	})
}
