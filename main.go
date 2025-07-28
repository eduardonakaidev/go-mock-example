// main.go
package main

import (
	"fmt"
	"github.com/eduardonakaidev/go-mock-example/repo"  // Altere para seu módulo real
	"github.com/eduardonakaidev/go-mock-example/service"
)

func main() {
	// Exemplo de uso real
	realRepo := repo.NewRealUserRepository()
	service := service.NewUserService(realRepo)
	
	user, err := service.GetUserName(1)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	
	fmt.Println("Usuário real:", user)
}