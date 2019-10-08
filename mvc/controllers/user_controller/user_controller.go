package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/fdiaz7/go-mricroservice/mvc/utils"

	"github.com/fdiaz7/go-mricroservice/mvc/services"
)

//GetUser controller
func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "User Id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		utils.RespondError(c, apiErr)
		return
	}
	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}
	//return user to client
	utils.Respond(c, http.StatusOK, user)
}
