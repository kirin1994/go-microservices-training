package domain

import (
	"fmt"
	"net/http"

	"github.com/kirin1994/go-microservices-training/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "Adrian", LastName: "Pod", Email: "apo@wp.pl"},
	}
)

func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exist", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
