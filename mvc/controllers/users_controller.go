package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shakilbd009/go-microsrvcs/mvc/services"
	"github.com/shakilbd009/go-microsrvcs/mvc/utils"
)

//GetUser returns
func GetUser(c *gin.Context) {

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		appErr := utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusNotFound,
			Code:       "bad_request",
		}
		utils.Respond(c, appErr.StatusCode, appErr)
		return
	}
	user, appErr := services.UsersService.GetUser(userID)
	if appErr != nil {
		utils.Respond(c, appErr.StatusCode, appErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)
}
