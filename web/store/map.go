package store

import "errors"

type storageMap struct {
	db map[string]string
}

func (s storageMap) SaveMapping(shortURL string, fullURL string) {
	s.db[shortURL] = fullURL
}

func (s storageMap) RetrieveURL(shortURL string) (string, error) {
	var err error
	v, ok := s.db[shortURL]
	if !ok {
		err = errors.New("short url not found")
	}
	return v, err
}

func (s storageMap) RetrieveAll() map[string]string {
	return s.db
}

func InitializeMap() storageMap {
	storage := storageMap{}
	storage.db = map[string]string{}
	return storage
}
