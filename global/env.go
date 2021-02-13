package global

import "os"

var (
	Port = GetEnv("SERVICING_PORT", "8080")
)

func GetEnv(env, fallback string) string {
	if v, ok := os.LookupEnv(env); ok {
		return v
	}

	return fallback
}