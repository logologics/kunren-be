package domain

import (
	"fmt"
	"strconv"
	"time"

	hsh "github.com/mitchellh/hashstructure"
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

// ToLanguage convets a string to a d.Language
func ToLanguage(lang string) Language {
	switch lang {
	case "ja": return Japanese
	default: return English
	} 
} 

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
	LemmaHash string `bson:"lhash"`
}

type Lemma struct {
	Reading  string    `json:"reading"`
	Lexeme   string    `json:"lexeme"`
	Key      string    `json:"key"`
	Meanings []Meaning `json:"meanings"`
}

// Hash calculates the Lemma's hash
func (l Lemma) Hash() (string, error) {
	h, err := hsh.Hash(l, nil)
	if err != nil {
		return "", fmt.Errorf("Could not calculate hash for %v", l.Key)
	}

	return strconv.FormatUint(h, 10), nil
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
	ID            mp.ObjectID `json:"_id" bson:"_id,omitempty"`
	Language      Language    `json:"language" bson:"language,omitempty"`
	WordID        mp.ObjectID `json:"wordID" bson:"wordID,omitempty"`
	UserID        mp.ObjectID `json:"userID" bson:"userID,omitempty"`
	SearchStrings []string    `json:"searchString" bson:"searchStrings,omitempty"`
	DateCreated   time.Time   `json:"dateCreated" bson:"dateCreated,omitempty"`
	DateSeen      time.Time   `json:"dateSeen" bson:"dateSeen,omitempty"`
	Seen          int         `json:"seen" bson:"seen,omitempty"`
	Confidence    int         `json:"confidence" bson:"confidence,omitempty"`
	Key           string      `json:"key" bson:"key,omitempty"`
	Tags					[]string		`json:"tags" bson:"tags,omitempty"`
}

type SearchResult struct {
	Query string `json:"query"`
	Words []Word `json:"words"`
}

type VocabListItem struct {
	ID            mp.ObjectID `json:"_id" bson:"_id,omitempty"`
	Language      Language    `json:"language" bson:"language,omitempty"`
	WordID        mp.ObjectID `json:"wordID" bson:"wordID,omitempty"`
	UserID        mp.ObjectID `json:"userID" bson:"userID,omitempty"`
	SearchStrings []string    `json:"searchString" bson:"searchStrings,omitempty"`
	DateCreated   time.Time   `json:"dateCreated" bson:"dateCreated,omitempty"`
	DateSeen      time.Time   `json:"dateSeen" bson:"dateSeen,omitempty"`
	Seen          int         `json:"seen" bson:"seen,omitempty"`
	Confidence    int         `json:"confidence" bson:"confidence,omitempty"`
	Key           string      `json:"key" bson:"key,omitempty"`
	Word *Word `json:"word" bson:"word,omitempty"`
}

type VocabPage struct {
	Vocabs []VocabListItem`json:"vocabs"`
	Seq int `json:"seq"`
	Size int `json:"size"`
	Count int `json:"cnt"`
	TotalCount int64 `json:"total"`
	IsLast bool `json:"isLast"`
	IsFirst bool `json:"isFirst"`
	Last int `json:"last"`
}

// Message is a simply message object to send in 
// some responses
type Message struct {
	Status int `json:"status"`
	Message string `json:"msg"`
}