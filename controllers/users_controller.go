package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kirin1994/go-microservices-training/utils"

	"github.com/kirin1994/go-microservices-training/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiError)
		resp.WriteHeader(apiError.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiError := services.GetUser(userID)
	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		resp.WriteHeader(apiError.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
