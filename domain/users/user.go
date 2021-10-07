package users

type User struct {
	id          int64  `json:"id"`
	firstName   string `json:"first_name"`
	lastName    string `json:"last_name"`
	email       string `json:"email"`
	createdDate string `json:"created_date"`
}
