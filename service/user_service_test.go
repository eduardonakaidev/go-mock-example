// service/user_service_test.go
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

	// 1. Criar o mock
	mockRepo := mock.NewMockUserRepository(ctrl)

	// 2. Configurar expectativas
	mockRepo.EXPECT().
		GetUser(123).
		Return(&repo.User{ID: 123, Name: "John Mock"}, nil)

	// 3. Injectar o mock
	userService := service.NewUserService(mockRepo)

	// 4. Executar o teste
	name, err := userService.GetUserName(123)
	
	// 5. Validações
	if err != nil {
		t.Fatalf("Erro inesperado: %v", err)
	}
	
	if name != "John Mock" {
		t.Errorf("Nome esperado: John Mock, obtido: %s", name)
	}
}