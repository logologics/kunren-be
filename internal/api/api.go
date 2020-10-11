package api

import (
	"net/http"

	d "github.com/logologics/kunren-be/internal/domain"
	r "github.com/logologics/kunren-be/internal/repo"
	mongo "github.com/logologics/kunren-be/internal/repo/mongo"
	log "github.com/sirupsen/logrus"
)

// Env is the api env
type Env struct {
	Config *d.Config
	Repo   r.Repo
}

// AppHandlerFunc that return error
type AppHandlerFunc func(http.ResponseWriter, *http.Request) error

// ServeHTTP calls
func (fn AppHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		switch v := err.(type) {
		case HTTPError:
			log.WithFields(log.Fields{"loc": "ServeHttp", "err": v, "ctx": v.Context}).Error(v.Message)
			v.SendError(w)
		default:
			log.WithFields(log.Fields{"loc": "ServeHttp", "err": err}).Error("Unexpected error")
			http.Error(w, "Unexpected server error", http.StatusInternalServerError)
		}
	}
}

// CreateRepo creates a new repo (only mongo supported)
func CreateRepo(c *d.Config) (r.Repo, error) {
	return mongo.Connect(c)
}
