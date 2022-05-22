package env

import (
	"log"
	"os"
)

func GetTokenStorePath() string {
	return getEnv("TOKEN_STORE_PATH", "Path")
}

func GetRegion() string {
	return getEnv("TOKEN_STORE_REGION", "Region")
}

func GetBucket() string {
	return getEnv("TOKEN_STORE_BUCKET", "Bucket")
}

func GetS3StoreKey() string {
	return getEnv("TOKEN_STORE_S3KEY", "S3 Key")
}

func getEnv(key string, visibleName string) string {
	env := os.Getenv(key)
	if len(env) == 0 {
		log.Fatalf("Environment (%v) not found\n", visibleName)
	}

	return env
}
