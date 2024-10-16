// Package storage/storage.go
package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Storage хранилище данных
type Storage struct {
	jsonPath string
	urls     []*urlData
	nextID   int
}

type urlData struct {
	Uuid        string `json:"uuid"`
	ShortUrl    string `json:"short_url"`
	OriginalUrl string `json:"original_url"`
}

// NewStorage создаёт storage, который будет хранить данные
// в json-файле по пути jsonPath.
func NewStorage(jsonPath string) *Storage {
	return &Storage{
		jsonPath: jsonPath,
	}
}

func (s *Storage) Load() error {
	bytes, err := os.ReadFile(s.jsonPath)
	if err != nil {
		return fmt.Errorf("ReadFile failed: %v", err)
	}
	err = json.Unmarshal(bytes, &s.urls)
	if err != nil {
		return fmt.Errorf("Unmarshal failed: %v", err)
	}
	max := 0
	for _, u := range s.urls {
		// ignore error here
		uuid, _ := strconv.Atoi(u.Uuid)
		if uuid > max {
			max = uuid
		}
	}
	s.nextID = max + 1
	return nil
}

func (s *Storage) SaveURL(randomCode string, origURL string) error {
	s.urls = append(s.urls, &urlData{
		Uuid:        strconv.Itoa(s.nextID),
		ShortUrl:    randomCode,
		OriginalUrl: origURL,
	})
	s.nextID++
	bytes, err := json.Marshal(s.urls)
	if err != nil {
		return fmt.Errorf("failed Marshal: %v", err)
	}
	err = os.WriteFile(s.jsonPath, bytes, 0666)
	if err != nil {
		return fmt.Errorf("failed Write File: %v", err)
	}
	return nil
}

func (s *Storage) LoadURL(code string) (string, error) {
	for _, u := range s.urls {
		if u.ShortUrl == code {
			return u.OriginalUrl, nil
		}
	}
	return "", fmt.Errorf("no url found by code %v", code)
}
