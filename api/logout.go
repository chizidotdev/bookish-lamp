package api

import (
	"log"
	"net/http"
	"net/url"

	"github.com/chizidotdev/copia/utils"
	"github.com/gin-gonic/gin"
)

// Handler for our logout.
func (server *Server) logout(ctx *gin.Context) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading config", err)
	}

	logoutUrl, err := url.Parse("https://" + config.Auth0Domain + "/v2/logout")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}

	// returnTo, err := url.Parse(scheme + "://" + ctx.Request.Host)
	returnTo, err := url.Parse(scheme + "://" + "localhost:3000/items")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", config.Auth0ClientID)
	logoutUrl.RawQuery = parameters.Encode()

	ctx.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}
