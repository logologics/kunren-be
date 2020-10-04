package crud

import (
	d "github.com/logologics/kunren-be/internal/domain"
)

func (db *Db) StoreUser(d.User) error {
	return nil
}

func (db *Db) LoadUser(id int) (d.User, error) {
	return d.User{}, nil
}

func (db *Db) UpdateUser(d.User) error {
	return nil
}

func (db *Db) DeleteUser(id int) error {
	return nil
}
