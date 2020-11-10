package route

/*
// Route represents a REST route
type Route struct {
	Name        string
	Method      string
	Path        string
	Queries     []string
	HandlerFunc api.AppHandlerFunc
	AcceptCT    string
}

func frontendHandler(env *api.Env) api.AppHandlerFunc {
	restEnv := rest.Env(*env)

	return restEnv.Index
}

func makeRoutes(env *api.Env) []Route {

	restEnv := rest.Env(*env)

	return []Route{
		{
			"RandomQuestions",
			"GET",
			"/api/rq", // TODO request paras
			[]string{},
			restEnv.AuthenticatedMW(restEnv.GenerateRandomQuestions),
			"application/json",
		},
		{
			"Search Jisho",
			"GET",
			"/api/search/jisho",
			[]string{"q", "{query}"},
			restEnv.AuthenticatedMW(restEnv.SearchJisho),
			"application/json",
		},
		{
			"Remember Word",
			"POST",
			"/api/remember",
			[]string{"t", "{tags}"},
			restEnv.AuthenticatedMW(restEnv.CheckRepo(restEnv.Remember)),
			"application/json",
		},
		{
			"Remember Word",
			"POST",
			"/api/remember",
			[]string{},
			restEnv.AuthenticatedMW(restEnv.CheckRepo(restEnv.Remember)),
			"application/json",
		},
		{
			"Get Vocabs",
			"GET",
			"/api/vocabs",
			[]string{"p", "{page:\\d+}", "s", "{sorting}", "ps", "{pageSize:\\d+}", "t", "{tags}"},
			restEnv.AuthenticatedMW(restEnv.CheckRepo(restEnv.Vocabs)),
			"",
		},
		{
			"Find Vocab (but only check)",
			"GET",
			"/api/vocabs/find",
			[]string{"k", "{key}", "l", "{lang}", "c", "{check:true|false}"},
			restEnv.AuthenticatedMW(restEnv.CheckRepo(restEnv.FindVocab)),
			"",
		},
		{
			"Find Vocab",
			"GET",
			"/api/vocabs/find",
			[]string{"k", "{key}", "l", "{lang}"},
			restEnv.AuthenticatedMW(restEnv.CheckRepo(restEnv.FindVocab)),
			"",
		},
		{
			"Get Tags",
			"GET",
			"/api/vocabs/tags",
			[]string{"l", "{lang}"},
			restEnv.AuthenticatedMW(restEnv.CheckRepo(restEnv.Tags)),
			"",
		},
		{
			"Delete Tag",
			"DELETE",
			"/api/vocabs/tags/{tag}",
			[]string{"l", "{lang}"},
			restEnv.AuthenticatedMW(restEnv.CheckRepo(restEnv.DeleteTag)),
			"",
		},

		{
			"Auth callback",
			"GET",
			"/auth/callback/{provider}",
			[]string{},
			restEnv.CheckRepo(restEnv.Callback),
			"",
		},
		{
			"Auth Logout",
			"GET",
			"/auth/logout/{provider}",
			[]string{},
			restEnv.CheckRepo(restEnv.Logout),
			"",
		},
		{
			"Authorize",
			"GET",
			"/auth/authorize/{provider}",
			[]string{},
			restEnv.CheckRepo(restEnv.Authorize),
			"",
		},
		{
			"Session info",
			"GET",
			"/auth/session",
			[]string{},
			restEnv.AuthenticatedMW(restEnv.CheckRepo(restEnv.Session)),
			"",
		},
	}
}

*/
