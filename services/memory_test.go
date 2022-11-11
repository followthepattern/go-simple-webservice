package services

import (
	"testing"

	mocks "github.com/followthepattern/go-simple-webservice/mocks/services"
	"github.com/golang/mock/gomock"
)

func TestMemory(t *testing.T) {
	db := mocks.NewMockDatabase(gomock.NewController(t))

	db.EXPECT().GetUserByID("u1")

	db.GetUserByID("u1")
}
