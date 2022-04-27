package utils

import "os"

func SetEnv(key, value string) error {
	return os.Setenv(key, value)
}

func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
