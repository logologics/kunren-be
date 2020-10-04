package domain

import (
	"time"
)
// Welcome contains the Kunren welcome message
type Welcome struct {
	Version string `json:"version"`
	Hello   string `json:"message"`
}

type ID string

type Language string
var Japanese Language = "ja"
var English Language = "en"

type Question struct {
	ID `json:"id"`
	Question string `json:"q"`
	Answer string `json:"a"`
	Features []string `json:"fs"`
}

type Questions struct {
	Questions []Question `json:"qs"`
}

type User struct {
	ID `json:"id"`
	Name string `json:"name"`
	Languages []Language `json:"languages"`
	Email string `json:"email"`
	LastLogin time.Time `json:"lastLogin"`
}

type Word struct {
	ID `json:"id"`
	Key string `json:"key"`
	Language Language `json:"language"`
	Source string `json:"src"`
	DateCreated time.Time `json:"dateCreated"`
	Lemma
}
	

type Lemma struct {
	ID `json:"id"`
	Reading string `json:"reading"`
	Lexeme string `json:"lexeme"`
	Key string `json:"key"`
	Meanings []Meaning `json:"meanings"`
}

type Meaning struct {
	ID `json:"id"`
	POS []string `json:"pos"`
	Translations []Translation `json:"translations"`
}

type Translation struct {
	ID `json:"id"`
	Language Language `json:"language"`
	Text string `json:"text"`
}

type Vocab struct {
	ID `json:"id"`
	Language Language `json:"language"`
	WordID ID `json:"wordID"`
	UserID ID `json:"useID"`
	SearchStrings []string `json:"searchString"`
	DateCreated time.Time `json:"dateCreated"`
	DateSeen time.Time `json:"dateSeen"`
	Seen int `json:"seen"`
	Confidence int `json:"confidence"`

}

type SearchResult struct {
	Query string `json:"query"`
	Words[] Word `json:"words"`
}