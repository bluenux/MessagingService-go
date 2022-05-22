package env

import (
	"log"
	"os"
)

func IsAWSLambda() bool {
	env, isAWS := os.LookupEnv("AWS_EXECUTION_ENV")
	log.Printf("env %v\n", env)
	log.Printf("is AWS ? : %v", isAWS)
	return isAWS
}
