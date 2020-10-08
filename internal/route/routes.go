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
		{
			"Index",
			"GET",
			"/",
			restEnv.Index,
			"",
		},
		{
			"RandomQuestions",
			"GET",
			"/rq", // TODO request paras
			restEnv.GenerateRandomQuestions,
			"application/json",
		},
		{
			"Search Jisho",
			"GET",
			"/search/jisho/{query}",
			restEnv.SearchJisho,
			"application/json",
		},
		{
			"Remember Word",
			"POST",
			"/remember",
			restEnv.Remember,
			"application/json",
		},
	}

}
