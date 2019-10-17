package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
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

func SaveStore(pathString string, store *Store) {
	storeByteData, err := json.MarshalIndent(*store, "", "\t")
	sigolo.FatalCheck(err)

	// Write copy with time stamp
	dir, _ := path.Split(pathString)
	err = ioutil.WriteFile(dir+"_"+time.Now().Format("2006-01-02")+".json", storeByteData, 0644)
	sigolo.FatalCheck(err)

	// (Over) write actual store
	err = ioutil.WriteFile(pathString, storeByteData, 0644)
	sigolo.FatalCheck(err)
}
