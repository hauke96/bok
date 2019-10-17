package main

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/hauke96/sigolo"
)

type Store struct {
	Entries []*Entry
}

type Entry struct {
	Amount int       `json:"amount"`
	Note   string    `json:"note"`
	Time   time.Time `json:"time"`
}

func ReadStore(path string) *Store {
	file, err := ioutil.ReadFile(path)
	sigolo.FatalCheck(err)

	var store Store
	json.Unmarshal(file, &store)

	return &store
}

func SaveStore(path string, store *Store) {
	file, err := json.MarshalIndent(*store, "", "\t")
	sigolo.FatalCheck(err)

	err = ioutil.WriteFile(path, file, 0644)
	sigolo.FatalCheck(err)
}
