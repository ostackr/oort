package environment

import (
	"os"
	"strconv"
)

func Get(key string) string {
	return os.Getenv(key)
}

func Find(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

func GetAsInt(key string) int {
	value := os.Getenv(key)
	integer, _ := strconv.Atoi(value)
	return integer
}

func FindAsInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	integer, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return integer
}
