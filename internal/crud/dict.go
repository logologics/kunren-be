package crud

import (
	d "github.com/logologics/kunren-be/internal/domain"
)

// StoreWord bla
func (db *Db) StoreWord(w d.Word) error {
	return nil
}

// LoadWord bla
func (db *Db) LoadWord(id int) (d.Word, error) {
	return d.Word{}, nil
}

// StoreWord bla
func (db *Db) UpdateWord(w d.Word) error {
	return nil
}

// StoreWord bla
func (db *Db) DeleteWord(id int) error {
	return nil
}
