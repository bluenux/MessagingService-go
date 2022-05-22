package store

import (
	"MessagingService/api/env"
	"log"
)

type TokenStore interface {
	Set(token string) bool
	All() []string
}

type tokenStore struct {
	tokenStore TokenStore
}

func (t tokenStore) Set(token string) bool {
	return t.tokenStore.Set(token)
}

func (t tokenStore) All() []string {
	return t.tokenStore.All()
}

func NewStore() TokenStore {
	if env.IsAWSLambda() {
		log.Println("Using AWS S3 store!")
		return &tokenStore{S3Store{
			fileStore:  FileStore{filePath: env.GetTokenStorePath()},
			region:     env.GetRegion(),
			bucket:     env.GetBucket(),
			s3StoreKey: env.GetS3StoreKey(),
		}}
	}
	log.Println("Using File store!")
	return &tokenStore{FileStore{filePath: env.GetTokenStorePath()}}
}
