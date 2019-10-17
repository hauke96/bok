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
	Entries []*Entry `json:"entries"`
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
