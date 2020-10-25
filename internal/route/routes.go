package route

import (
	"github.com/logologics/kunren-be/internal/api"
	"github.com/logologics/kunren-be/internal/rest"
)

// Route represents a REST route
type Route struct {
	Name        string
	Method      string
	Path        string
	Queries     []string
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
			[]string{},
			restEnv.Index,
			"",
		},
		{
			"RandomQuestions",
			"GET",
			"/rq", // TODO request paras
			[]string{},
			restEnv.GenerateRandomQuestions,
			"application/json",
		},
		{
			"Search Jisho",
			"GET",
			"/search/jisho",
			[]string{"q", "{query}"},
			restEnv.SearchJisho,
			"application/json",
		},
		{
			"Remember Word",
			"POST",
			"/remember",
			[]string{},
			restEnv.CheckRepo(restEnv.Remember),
			"application/json",
		},
		{
			"Get Vocabs",
			"GET",
			"/vocabs",
			[]string{"k", "{key}"},
			restEnv.CheckRepo(restEnv.Vocabs),
			"",
		},
		{
			"Find Vocab (but only check)",
			"GET",
			"/vocabs/find",
			[]string{"k", "{key}", "l", "{lang}", "c", "{check:true|false}"},
			restEnv.CheckRepo(restEnv.FindVocab),
			"",
		},
		{
			"Find Vocab",
			"GET",
			"/vocabs/find",
			[]string{"k", "{key}", "l", "{lang}"},
			restEnv.CheckRepo(restEnv.FindVocab),
			"",
		},
	}
}
