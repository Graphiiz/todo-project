package users

type User struct {
	// Id    string   `json:"id"`
	// Name  string   `json:"name"`
	Email string   `json:"email"`
	Todos []string `json:"todos"`
}
