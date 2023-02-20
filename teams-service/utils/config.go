package helpers

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	AuthService          string        `mapstructure:"AUTH_SERVICE_URL"`
	Login             	 string        `mapstructure:"AUTH_LOGIN"`
	AuthUsername         string        `mapstructure:"AUTH_USERNAME"`
	AuthPassword         string        `mapstructure:"AUTH_PASSWORD"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// LoadMockConfig returns a mock configuration for testing purposes.
func LoadMockConfig() (config Config, err error) {
	config = Config{
		DBDriver:             "postgres",
		DBSource:             "postgres://postgres:password@localhost:5434/team?sslmode=disable",
	}
	return
}