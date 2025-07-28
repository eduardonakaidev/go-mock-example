package service_test

import (
	"testing"
	mock "github.com/eduardonakaidev/go-mock-example/mocks"
	"github.com/eduardonakaidev/go-mock-example/repo"
	"github.com/eduardonakaidev/go-mock-example/service"
	"go.uber.org/mock/gomock"
)

func TestGetUserName(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := mock.NewMockUserRepository(ctrl)
    mockRepo.EXPECT().
        GetUser(123).
        Return(&repo.User{ID: 123, Name: "John Mock"}, nil)

    userService := service.NewUserService(mockRepo)
    name, err := userService.GetUserName(123)
    
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    
    if name != "John Mock" {
        t.Errorf("Expected: John Mock, Got: %s", name)
    }
}