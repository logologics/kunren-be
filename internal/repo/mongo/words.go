package mongo

import (
	d "github.com/logologics/kunren-be/internal/domain"
)

// StoreWord bla
func (mongo *Mongo) StoreWord(w d.Word) (d.ID, error) {
	return "", nil
}

// LoadWord bla
func (mongo *Mongo) LoadWord(id d.ID) (d.Word, error) {
	return d.Word{}, nil
}

// UpdateWord bla
func (mongo *Mongo) UpdateWord(w d.Word) error {
	return nil
}

// DeleteWord bla
func (mongo *Mongo) DeleteWord(id d.ID) error {
	return nil
}

// ListWords lists all words in the dict
func (mongo *Mongo) ListWords() ([]d.Word, error){
	return nil, nil
}
