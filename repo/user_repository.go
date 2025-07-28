package repo

type User struct {
    ID   int
    Name string
}

type UserRepository interface {
    GetUser(id int) (*User, error)
}

type realUserRepository struct{}

func NewRealUserRepository() UserRepository {
    return &realUserRepository{}
}

func (r *realUserRepository) GetUser(id int) (*User, error) {
    return &User{ID: id, Name: "Real User"}, nil
}