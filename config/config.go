package config

import (
	"github.com/spf13/viper"
)

// Config struct that contain all config for this application
type Config struct {
	Application ApplicationConfig
	Server      ServerConfig
	Database    DatabaseConfig
}

// Application config
type ApplicationConfig struct {
	// logger level (all possible value: panic, fatal, error, warn, warning, info, debug and trace)
	LogLevel string
}

// HTTP server config
type ServerConfig struct {
	// HTTP port
	Port string
}

// Database config
type DatabaseConfig struct {
	// database type (possible to switching from mysql to postgresql/mongodb/memory etc)
	Type string
	// username
	User string
	// password
	Password string
	// database address
	Addr string
	// database name
	DBName string
}

// get string config from environment
// if value is not found fomr environemnt, defaultValue will be used
func getStringConfigWithDefault(key, defaultValue string) (string, error) {
	err := viper.BindEnv(key, key)
	if err != nil {
		return "", err
	}

	value := viper.GetString(key)
	if value == "" {
		return defaultValue, nil
	}

	return value, nil
}

// Loading all confing from environment
// Using os environment due to this application is expected to be deployed to docker/k8s
// setting environment is easiest way to config application in docker/k8s comparing reading config file
func LoadConfig() (*Config, error) {
	c := &Config{}
	var err error

	// Application Config
	c.Application.LogLevel, err = getStringConfigWithDefault("LOG_LEVEL", "info")
	if err != nil {
		return nil, err
	}

	// Server Config
	c.Server.Port, err = getStringConfigWithDefault("SERVER_PORT", "3000")
	if err != nil {
		return nil, err
	}

	// Database Config
	c.Database.Type, err = getStringConfigWithDefault("DB_TYPE", "mysql")
	if err != nil {
		return nil, err
	}

	c.Database.User, err = getStringConfigWithDefault("DB_USER", "user")
	if err != nil {
		return nil, err
	}

	c.Database.Password, err = getStringConfigWithDefault("DB_PASSWORD", "password")
	if err != nil {
		return nil, err
	}

	c.Database.Addr, err = getStringConfigWithDefault("DB_ADDRESS", "localhost:3306")
	if err != nil {
		return nil, err
	}

	c.Database.DBName, err = getStringConfigWithDefault("DB_NAME", "getground")
	if err != nil {
		return nil, err
	}

	return c, nil
}
