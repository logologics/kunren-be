package rest

import (
	"fmt"
	"net/http"

	"github.com/logologics/kunren-be/internal/api"
)

func (e *Env) CheckRepo(next api.AppHandlerFunc) api.AppHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) error {
		if e.Repo.Ready() {
			return next(w, r)
		}
		err := fmt.Errorf("repo is not ready")
		return api.NewHTTPInternalServerError(err, "Something went wrong", "rest searchJisho()")

	}

}
