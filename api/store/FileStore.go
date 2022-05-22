package store

import (
	"bufio"
	"log"
	"os"
)

type FileStore struct {
}

func (f FileStore) All() []string {
	file, err := os.Open("/mnt/data/tokens")
	if err != nil {
		log.Printf("error : %v\n", err)
	}
	defer file.Close()

	var list []string

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		text := scanner.Text()
		log.Printf("data : %v\n", text)
		list = append(list, text)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("error : %v\n", err)
	}

	return list
}

func (f FileStore) Set(token string) bool {
	if f.isNewToken(token) {
		log.Printf("store token value...")
		fs, err := os.OpenFile("/mnt/data/tokens", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Printf("error : %v\n", err)
		}

		n, err := fs.WriteString(token)
		log.Printf("store token result %v", n)
		if err != nil {
			log.Printf("error : %v\n", err)
		}

		n1, err := fs.WriteString("\n")
		log.Printf("store token result.. %v", n1)
		if err != nil {
			log.Printf("error : %v\n", err)
		}

		err = fs.Close()
		if err != nil {
			log.Printf("error : %v\n", err)
		}

		return true
	}
	return false
}

func (f FileStore) isNewToken(token string) bool {
	for _, element := range f.All() {
		if token == element {
			log.Printf("already registed!! : %v\n", token)
			return false
		}
	}
	return true
}
