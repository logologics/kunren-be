package route

import (
	"github.com/logologics/kunren-be/internal/api"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

// NewRestRouter creates a new REST router
func NewRestRouter(env *api.Env) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range makeRoutes(env) {

		r := router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
		if route.AcceptCT != "" {
			r.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
				return strings.HasPrefix(r.Header.Get("Accept"), route.AcceptCT) ||
					strings.HasPrefix(r.Header.Get("ContentType"), route.AcceptCT)
			})			
		}
		if len(route.Queries) > 0 {
			r.Queries(route.Queries...)
		}
	}

	return router
}
