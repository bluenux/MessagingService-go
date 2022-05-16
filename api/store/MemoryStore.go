package store

import "log"

type MemoryStore struct {
	deviceList []string
}

func (m MemoryStore) All() []string {
	return m.deviceList
}

func (m *MemoryStore) Set(token string) bool {
	if m.isNewToken(token) {
		m.deviceList = append(m.deviceList, token)
		log.Printf("added token : %v\n", token)

		return true
	}

	return false
}

func (m MemoryStore) isNewToken(token string) bool {
	for _, element := range m.deviceList {
		if token == element {
			log.Printf("already registed!! : %v\n", token)
			return false
		}
	}
	return true
}
