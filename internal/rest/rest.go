package rest

import (
	"encoding/json"
	"net/http"

	"github.com/logologics/kunren-be/internal/api"
	d "github.com/logologics/kunren-be/internal/domain"
)

// Env is a local env type
type Env api.Env

// GenerateRandomQuestions returns the hanlder for GET /GenerateRandomQuestions
func (e *Env) GenerateRandomQuestions(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	qs := d.Questions {
		Questions: map[string]d.Question{
		"q1": d.Question{
			Question: "q1",
			Answer:   "a1",
			Features: []string{"plain", "past", "conditional", "hallo", "good morning"},
		},
		"q2": d.Question{
			Question: "q2",
			Answer:   "aa",
			Features: []string{"plain", "past", "q2"},
		},
		"q3": d.Question{
			Question: "q3",
			Answer:   "a3",
			Features: []string{"polite", "present", "q3"},
		},
	},
	}

	return json.NewEncoder(w).Encode(qs)
}

// Index returns the hanlder for GET /
func (e *Env) Index(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	welcome := d.Welcome{Version: d.Version, Hello: "Wecome to Kunren!"}
	if err := json.NewEncoder(w).Encode(welcome); err != nil {
		return api.NewHttpBadRequest("Unexpected error in Index")
	}

	return nil
}
