package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/hauke96/sigolo"
)

type Store struct {
	Path    string   `json:"-"`
	Dirty   bool     `json:"-"` // "true" when entry set changed since last save
	Entries []*Entry `json:"entries"`
}

type Entry struct {
	Amount   int       `json:"amount"`
	Date     time.Time `json:"date"`
	Note     string    `json:"note"`
	Category string    `json:"category"`
}

// AddEntry will add an entry to the stores list of entries. It also sets the
// dirty-flag to "true".
func (s *Store) AddEntry(amount int, date time.Time, description string, category string) {
	var e Entry

	e.Amount = amount
	e.Note = description
	e.Date = date
	e.Category = category

	s.Entries = append(s.Entries, &e)
	s.Dirty = true

	sigolo.Debug("Added entry [%d, %s, '%s', '%s']", amount, date, description, category)
}

func (s *Store) HasEntry(amount int, date time.Time, description string, category string) bool {
	for _, e := range s.Entries {
		// Only use amount and date
		if e.Amount == amount &&
			e.Date == date {
			return true
		}
	}
	return false
}

// SaveStore will save the store to its location on disk. This will also set the
// dirty-flag back to "false".
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

func (s *Store) filterByDatePrefix(prefix string) *Store {
	var store = &Store{
		Path:    s.Path + "_filtered.json",
		Entries: make([]*Entry, 0),
	}

	for _, e := range s.Entries {
		if strings.HasPrefix(e.Date.Format("2006-01-02"), prefix) {
			store.AddEntry(e.Amount, e.Date, e.Note, e.Category)
		}
	}

	sort.Slice(store.Entries, func(i, j int) bool { return store.Entries[i].Date.Before(store.Entries[j].Date) })

	return store
}

// ReadStore reads the store from disk.
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
