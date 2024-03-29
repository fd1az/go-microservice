package oauth

import (
	"net/http"

	"github.com/fdiaz7/go-microservice/oauth-api/src/api/domain/oauth"
	"github.com/fdiaz7/go-microservice/oauth-api/src/api/services"
	"github.com/fdiaz7/go-microservice/src/api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, token)
}

func GetAccessToken(c *gin.Context) {
	tokenId := c.Param("token_id")
	token, err := services.OauthService.GetAccessToken(tokenId)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}
