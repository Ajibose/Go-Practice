package model

type Data struct {
	ID        string
	CreatedAt string
	Body      string
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Storage struct {
	Users []User
	Ids   map[string]Data
}
