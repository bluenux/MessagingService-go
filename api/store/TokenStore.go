package store

import (
	"MessagingService/api/utility"
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
	if utility.IsAWSLambda() {
		return &tokenStore{FileStore{}}
	}
	return &tokenStore{&MemoryStore{}}
}
