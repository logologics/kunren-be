package mongo

import (
	d "github.com/logologics/kunren-be/internal/domain"
)

// StoreUser bla
func (mongo *Mongo) StoreUser(d.User) (d.ID, error) {
	return "", nil
}

// LoadUser bla
func (mongo *Mongo) LoadUser(id d.ID) (d.User, error) {
	return d.User{}, nil
}

//UpdateUser bla
func (mongo *Mongo) UpdateUser(u d.User) error {
	return nil
}

// DeleteUser bla
func (mongo *Mongo) DeleteUser(id d.ID) error {
	return nil
}
