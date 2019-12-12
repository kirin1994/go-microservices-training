package services

import (
	"github.com/kirin1994/go-microservices-training/domain"
	"github.com/kirin1994/go-microservices-training/utils"
)

func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
