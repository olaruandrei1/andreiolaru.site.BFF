package config

import (
	"os"
	"strings"
)

func GetEnvAsSlice(key string, sep string) []string {
	val := os.Getenv(key)
	if val == "" {
		return []string{}
	}
	return strings.Split(val, sep)
}
