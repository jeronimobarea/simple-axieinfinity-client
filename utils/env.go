package utils

import "os"

func GetEnvFallback(name, fallback string) string {
	if res, ok := os.LookupEnv(name); ok {
		return res
	}
	return fallback
}
