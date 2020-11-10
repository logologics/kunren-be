package route

import (
	"github.com/gin-gonic/gin"
	"github.com/logologics/kunren-be/internal/api"
	"github.com/logologics/kunren-be/internal/domain"
)

// New creates a new gin router
func New(c *domain.Config) (*gin.Engine, error) {

	env, err := api.CreateEnv(c)
	if err != nil {
		return nil, err
	}

	router := gin.New()
	router.Use(env.CheckRepo())

	// api
	api := router.Group("/api")
	api.Use(env.Authenticated())
	api.GET("/rq", env.GenerateRandomQuestions)
	api.GET("/search/jisho", env.SearchJisho)
	api.POST("/remember", env.Remember)
	api.GET("/vocabs", env.Vocabs)
	api.GET("/vocabs/find", env.FindVocab)
	api.GET("/vocabs/tags", env.Tags)
	api.DELETE("/vocabs/tags/:tag", env.DeleteTag)

	// auth
	auth := router.Group("/auth")
	auth.GET("/callback/:provider", env.Callback)
	auth.GET("/authorize/:provider", env.Authorize)
	auth.GET("/logout/:provider", env.Logout).Use(env.Authenticated())

	// auth authenticated
	session := auth.Group("/session")
	session.Use(env.Authenticated())
	session.GET("/", env.Session).Use(env.Authenticated())

	// frontend
	router.Static("/static", c.FrontEndPath+"/static")
	router.StaticFile("/asset-manifest.json", c.FrontEndPath+"/asset-manifest.json")
	router.StaticFile("/index.html", c.FrontEndPath+"/index.html")
	router.StaticFile("/kunren.png", c.FrontEndPath+"/kunren.png")
	router.StaticFile("/manifest.json", c.FrontEndPath+"/manifest.json")
	router.StaticFile("/service-worker.json", c.FrontEndPath+"/service-worker.json")
	router.StaticFile("/", c.FrontEndPath+"/index.html")

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found: " + ctx.Request.URL.Path})
	})
	return router, nil
}
