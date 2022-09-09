package env

import (
	"os"
)

// Get returns the value of the environment variable named by the key or the fallback value if the variable is not set.
func Get(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}

	return fallback
}

// Set sets the value of the environment variable named by the key.
func Set(key string, value string) {
	os.Setenv(key, value)
}
