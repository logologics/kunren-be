package api

import (
	"fmt"
	"net/http"

	d "github.com/logologics/kunren-be/internal/domain"
	r "github.com/logologics/kunren-be/internal/repo"
	mongo "github.com/logologics/kunren-be/internal/repo/mongo"
	log "github.com/sirupsen/logrus"
)

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
		case HttpError:
			v.sendError(w)
		default:
			errMsg := fmt.Sprintf("Unexpected error: %v", err)
			log.WithFields(log.Fields{"loc": "ServeHttp", "msg": errMsg})
			http.Error(w, errMsg, http.StatusInternalServerError)
		}
	}
}

func CreateRepo(config *d.Config) (r.Repo, error) {
	return &mongo.Mongo{}, nil
}
