package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

// Authenticated is a middleware that check that
// user has a valid session
func (e *Env) Authenticated() gin.HandlerFunc {

	return func(c *gin.Context) {

		email, err := gothic.GetFromSession(kunrenUserKey, c.Request)
		if err != nil {
			sendError(c, http.StatusUnauthorized, err, "Not authorized")
			return
		}

		if email == "" {
			err := fmt.Errorf("Email empty")
			sendError(c, http.StatusUnauthorized, err, "Not authorized")
			return
		}

		c.Next()

	}
}

// CheckRepo is a middleware that ensures the repository
// has been initialized
func (e *Env) CheckRepo() gin.HandlerFunc {

	return func(c *gin.Context) {
		if e.Repo.Ready() {
			c.Next()
			return
		}

		err := fmt.Errorf("Unexpected error")
		sendError(c, http.StatusInternalServerError, err, "Repo not initialized")
	}

}
