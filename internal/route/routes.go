package route

import (
	"github.com/logologics/kunren-be/internal/api"
	"github.com/logologics/kunren-be/internal/rest"
)

// Route represents a REST route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc api.AppHandlerFunc
	AcceptCT    string
}

func makeRoutes(env *api.Env) []Route {

	restEnv := rest.Env(*env)

	return []Route{
		Route{
			"Index",
			"GET",
			"/",
			restEnv.Index,
			"",
		},
		Route{
			"RandomQuestions",
			"GET",
			"/rq", // TODO request paras
			restEnv.GenerateRandomQuestions,
			"application/json",
		},
	}
}
