package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
)

func checkKey(key string) {
	if !viper.IsSet(key) && os.Getenv(key) == "" {
		log.Fatalf("panic: %s key not set", key)
	}
}

func getStringKey(key string) string {
	checkKey(key)
	return viper.GetString(key)
}

func getIntKey(key string) int {
	checkKey(key)
	value, err := strconv.Atoi(getStringKey(key))

	if err != nil {
		panicIfErrorForKey(err, key)
	}
	return value
}

func getBoolKey(key string) bool {
	checkKey(key)
	value, err := strconv.ParseBool(getStringKey(key))
	if err != nil {
		panicIfErrorForKey(err, key)
	}
	return value
}

func panicIfErrorForKey(err error, key string) {
	log.Fatal("Could not parse key: %s, Error: %s", key, err)
}

func loadConfigFromFile() {
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Config file not found")
	}
}
