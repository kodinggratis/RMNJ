package domain

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Interface sebagai kontrak antar layer
type UserService interface {
	GetByID(id int) (User, error)
}
