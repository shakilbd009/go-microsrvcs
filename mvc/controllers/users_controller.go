package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shakilbd009/go-microsrvcs/mvc/services"
	"github.com/shakilbd009/go-microsrvcs/mvc/utils"
)

//GetUser returns
func GetUser(resp http.ResponseWriter, req *http.Request) {

	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		appErr := utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusNotFound,
			Code:       "bad_request",
		}
		jsonData, _ := json.MarshalIndent(appErr, "", "  ")
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(appErr.StatusCode)
		resp.Write(jsonData)
		return
	}
	user, appErr := services.GetUser(userID)
	if appErr != nil {
		jsonData, _ := json.MarshalIndent(appErr, "", "  ")
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(appErr.StatusCode)
		resp.Write(jsonData)
		return
	}
	jsonData, _ := json.MarshalIndent(user, "", "  ")
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(200)
	resp.Write(jsonData)
}
