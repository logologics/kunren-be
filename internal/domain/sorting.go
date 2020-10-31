package domain

import (
	"fmt"
	"strconv"
	"strings"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// SortOrder encodes the sort order
type SortOrder int

const (
	// ASC encodes ascending sort order
	ASC SortOrder = 1
	// DESC encodes descending sort order
	DESC = -1
)

// Field is the name of a field to be sorted
type Field string

// Sorting defined the sorting for one data field
type Sorting struct {
	Field
	SortOrder
}

// ParseSorting parses a string in the format
// FIELDNAME:ORDER:FIELDNAME:ORDER
// into a sorting
func ParseSorting(s string) ([]Sorting, error) {
	if s == "" {
		return []Sorting{}, nil
	}
	parts := strings.Split(s, ":")
	l := len(parts)
	if l%2 != 0 {
		return []Sorting{}, fmt.Errorf("Invalid sorting string %v", s)
	}
	fs := []Sorting{}
	for i := 0; i < l; i = i + 2 {
		o, err := strconv.Atoi(parts[i+1])
		if err != nil {
			return []Sorting{}, err
		}
		fs = append(fs, Sorting{Field: Field(parts[i]), SortOrder: SortOrder(o)})
	}

	return fs, nil
}

func SortingToBson(srt []Sorting) bson.M {
	m := bson.M{}
	for _, s := range srt {
		m[string(s.Field)] = s.SortOrder
	}
	log.Infof("Sorting %v",  m)
	return m
}
