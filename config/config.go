package config

type appConfig struct {
	name          string
	port          int
	logevel       string
	database      DatabaseConfig
	sentryEnabled bool
}

var config appConfig

func Load() {
	loadConfigFromFile()
	config = appConfig{
		name:    getStringKey("APP_NAME"),
		port:    getIntKey("APP_PORT"),
		logevel: getStringKey("LOG_LEVEL"),
		database: DatabaseConfig{
			host:        getStringKey("DB_HOST"),
			name:        getStringKey("DB_NAME"),
			user:        getStringKey("DB_USER"),
			password:    getStringKey("DB_PASSWORD"),
			port:        getIntKey("DB_PORT"),
			maxIdleConn: getIntKey("DB_MAX_IDLE_CONNECTION"),
			maxPoolSize: getIntKey("DB_MAX_POOL_SIZE"),
		},
		sentryEnabled: getBoolKey("SENTRY_ENABLED"),
	}
}

func AppPort() int {
	return config.port
}

func AppName() string {
	return config.name
}

func LogLevel() string {
	return config.logevel
}

func Database() DatabaseConfig {
	return config.database
}

func SentryEnabled() bool {
	return config.sentryEnabled
}
