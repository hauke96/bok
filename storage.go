package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/hauke96/sigolo"
)

type Store struct {
	Path    string
	Entries []*Entry
}

type Entry struct {
	Amount int       `json:"amount"`
	Note   string    `json:"note"`
	Date   time.Time `json:"date"`
}

func (s *Store) AddEntry(amount int, description string, date time.Time) {
	var e Entry

	e.Amount = amount
	e.Note = description
	e.Date = date

	s.Entries = append(s.Entries, &e)

	sigolo.Debug("Added entry [%s, %s, '%s']", amount, description, date)
}

func ReadStore(path string) *Store {
	var store Store

	store.Path = path
	store.Entries = make([]*Entry, 0)

	// Return empty store if given file does not exist
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		file, err := ioutil.ReadFile(path)
		sigolo.FatalCheck(err)

		json.Unmarshal(file, &store)
	}

	return &store
}

func SaveStore(path string, store *Store) {
	file, err := json.MarshalIndent(*store, "", "\t")
	sigolo.FatalCheck(err)

	err = ioutil.WriteFile(path, file, 0644)
	sigolo.FatalCheck(err)
}
