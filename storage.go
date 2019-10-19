package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/hauke96/sigolo"
)

type Store struct {
	Path    string   `json:"-"`
	Dirty   bool     `json:"-"` // "true" when changes happened that are not stored to disk
	Entries []*Entry `json:"entries"`
}

type Entry struct {
	Amount   int       `json:"amount"`
	Date     time.Time `json:"date"`
	Note     string    `json:"note"`
	Category string    `json:"category"`
}

func (s *Store) AddEntry(amount int, date time.Time, description string, category string) {
	var e Entry

	e.Amount = amount
	e.Note = description
	e.Date = date
	e.Category = category

	s.Entries = append(s.Entries, &e)
	s.Dirty = true

	sigolo.Debug("Added entry [%s, %s, '%s', '%s']", amount, date, description, category)
}

func (s *Store) SaveStore() {
	storeByteData, err := json.MarshalIndent(*s, "", "\t")
	sigolo.FatalCheck(err)

	// Write copy with time stamp
	dir, file := path.Split(s.Path)
	fileName := strings.TrimSuffix(file, filepath.Ext(file))
	err = ioutil.WriteFile(dir+fileName+"_"+time.Now().Format("2006-01-02_15:04:05")+".json", storeByteData, 0644)
	sigolo.FatalCheck(err)

	// (Over) write actual store
	err = ioutil.WriteFile(s.Path, storeByteData, 0644)
	sigolo.FatalCheck(err)

	s.Dirty = false
}

func ReadStore(path string) *Store {
	var store = &Store{
		Path:    path,
		Entries: make([]*Entry, 0),
	}

	// Return empty store if given file does not exist
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		file, err := ioutil.ReadFile(path)
		sigolo.FatalCheck(err)

		json.Unmarshal(file, &store)
	}

	store.Dirty = false

	return store
}
