package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/logologics/kunren-be/internal/api"
	d "github.com/logologics/kunren-be/internal/domain"
	jisho "github.com/logologics/kunren-be/internal/extDict/jisho"
)

// Env is a local env type
type Env api.Env

var pageSize = 10

// SearchJisho returns the handler for GET /search/jisho/{query}
func (e *Env) SearchJisho(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	query := vars["query"]

	sr, err := jisho.Search(query)
	if err != nil {
		return api.NewHTTPInternalServerError(err, "Something went wrong", "rest searchJisho()")
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(sr)
}

// Vocabs returns all vocabs for the user
func (e *Env) Vocabs(w http.ResponseWriter, r *http.Request) error {
	key := mux.Vars(r)["key"]
	log.Infof("Key %v", key)

	vocabs, err := e.Repo.ListVocabs(key, pageSize, e.User)
	if err != nil {
		return api.NewHTTPInternalServerError(err, "Can't close body", "rest - Remember()")
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(vocabs)
}

// Remember stores a search result in the dict history
func (e *Env) Remember(w http.ResponseWriter, r *http.Request) error {
	word := d.Word{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return api.NewHTTPBadRequest(err, "Cant read body", "rest - Remember()")
	}

	if err := r.Body.Close(); err != nil {
		return api.NewHTTPInternalServerError(err, "Can't close body", "rest - Remember()")
	}

	if err := json.Unmarshal(body, &word); err != nil {
		return api.NewHTTPInternalServerError(err, "Can't unmarshal body", "rest - Remember()")
	}

	storedWord, err := e.Repo.StoreWord(word)
	if err != nil {
		return api.NewHTTPInternalServerError(err, "Can't store", "rest - Remember()")
	}

	vocab := d.Vocab{
		Key:           storedWord.Key,
		WordID:        storedWord.ID,
		UserID:        e.User.ID,
		Language:      word.Language,
		SearchStrings: []string{storedWord.Key, storedWord.Lemma.Reading},
	}
	_, err = e.Repo.StoreVocab(vocab, true)
	if err != nil {
		return api.NewHTTPInternalServerError(err, "Can't create vocab", "rest - Remember()")
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(storedWord.ID)
}

// GenerateRandomQuestions returns the handler for GET /GenerateRandomQuestions
func (e *Env) GenerateRandomQuestions(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Access-Control-Allow-Origin", e.Config.KunrenFe)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	qs := d.Questions{
		Questions: []d.Question{
			{
				ID:       "1",
				Question: "q1",
				Answer:   "a1",
				Features: []string{"plain", "past", "conditional", "hallo", "good morning"},
			},
			{
				ID:       "2",
				Question: "q2",
				Answer:   "aa",
				Features: []string{"plain", "past", "q2"},
			},
			{
				ID:       "3",
				Question: "q3",
				Answer:   "a3",
				Features: []string{"polite", "present", "q3"},
			},
		},
	}

	return json.NewEncoder(w).Encode(qs)
}

// Index returns the handler for GET /
func (e *Env) Index(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	welcome := d.Welcome{Version: d.Version, Hello: "Welcome to Kunren!"}
	if err := json.NewEncoder(w).Encode(welcome); err != nil {
		return api.NewHTTPBadRequest(err, "Unexpected error in Index", "rest - index()")
	}

	return nil
}
