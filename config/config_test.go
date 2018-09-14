package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {

	configVars := map[string]string{
		"APP_NAME":               "sample",
		"APP_PORT":               "8080",
		"SENTRY_ENABLED":         "true",
		"LOG_LEVEL":              "info",
		"DB_HOST":                "localhost",
		"DB_USER":                "test",
		"DB_PASSWORD":            "12345",
		"DB_NAME":                "test",
		"DB_PORT":                "5432",
		"DB_MAX_IDLE_CONNECTION": "10",
		"DB_MAX_POOL_SIZE":       "10",
	}

	for k, v := range configVars {
		os.Setenv(k, v)
		defer os.Unsetenv(k)
	}

	Load()
	assert.Equal(t, 8080, AppPort())
	assert.Equal(t, configVars["LOG_LEVEL"], LogLevel())

	expectedDbConfig := DatabaseConfig{
		host:        "localhost",
		port:        5432,
		name:        "test",
		user:        "test",
		password:    "12345",
		maxIdleConn: 10,
		maxPoolSize: 10,
	}
	assert.Equal(t, expectedDbConfig, Database())
	assert.Equal(t, true, SentryEnabled())
}
