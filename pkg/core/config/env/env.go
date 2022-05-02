package env

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	godotenv.Load(".env")
}

func Get(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("Environment Variable not set: " + key)
	}

	return value, nil
}

func GetMany(keys []string) (map[string]string, error) {
	results := make(map[string]string)

	for _, key := range keys {
		value, err := Get(key)
		if err != nil {
			return make(map[string]string), err
		}
		results[key] = value
	}

	return results, nil
}
