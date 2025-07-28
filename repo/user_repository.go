// repo/user_repository.go
package repo

// User representa um usuário no sistema
type User struct {
	ID   int
	Name string
}

// UserRepository define a interface de acesso a dados
type UserRepository interface {
	GetUser(id int) (*User, error)
}

// Implementação real (ex: banco de dados)
type realUserRepository struct{}

func NewRealUserRepository() UserRepository {
	return &realUserRepository{}
}

func (r *realUserRepository) GetUser(id int) (*User, error) {
	// Em implementação real, isso viria de um banco de dados
	return &User{ID: id, Name: "Usuário Real"}, nil
}