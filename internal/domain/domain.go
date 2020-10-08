package domain

import (
	"time"

	mp "go.mongodb.org/mongo-driver/bson/primitive"
)

// Welcome contains the Kunren welcome message
type Welcome struct {
	Version string `json:"version"`
	Hello   string `json:"message"`
}

type Language string

var Japanese Language = "ja"
var English Language = "en"

type Question struct {
	ID       string   `json:"id"`
	Question string   `json:"q"`
	Answer   string   `json:"a"`
	Features []string `json:"fs"`
}

type Questions struct {
	Questions []Question `json:"qs"`
}

type User struct {
	ID        mp.ObjectID `json:"id"`
	Name      string      `json:"name"`
	Languages []Language  `json:"languages"`
	Email     string      `json:"email"`
	LastLogin time.Time   `json:"lastLogin"`
}

type Word struct {
	ID          mp.ObjectID `json:"_id" bson:"_id,omitempty"`
	Key         string      `json:"key" bson:"key,omitempty"`
	Language    Language    `json:"language" bson:"language,omitempty"`
	Source      string      `json:"src" bson:"src,omitempty"`
	DateCreated time.Time   `json:"dateCreated" bson:"dateCreated,omitempty"`
	Lemma
}

type Lemma struct {
	Reading  string    `json:"reading"`
	Lexeme   string    `json:"lexeme"`
	Key      string    `json:"key"`
	Meanings []Meaning `json:"meanings"`
}

type Meaning struct {
	POS          []string      `json:"pos"`
	Translations []Translation `json:"translations"`
}

type Translation struct {
	Language Language `json:"language"`
	Text     string   `json:"text"`
}

type Vocab struct {
	ID            mp.ObjectID `json:"id"`
	Language      Language    `json:"language"`
	WordID        mp.ObjectID `json:"wordID"`
	UserID        mp.ObjectID `json:"useID"`
	SearchStrings []string    `json:"searchString"`
	DateCreated   time.Time   `json:"dateCreated"`
	DateSeen      time.Time   `json:"dateSeen"`
	Seen          int         `json:"seen"`
	Confidence    int         `json:"confidence"`
}

type SearchResult struct {
	Query string `json:"query"`
	Words []Word `json:"words"`
}
