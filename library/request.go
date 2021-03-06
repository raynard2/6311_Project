package library

type CreateUserParams struct {
	Email           string `json:"email"`
	FullName        string `json:"full_name"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Method          string `json:"method" bson:"method"`
}

type LoginParams struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required, email"`
	Method   string `json:"method" bson:"method"`
}

type SetPasswordParams struct {
	Hash     string `json:"hash" validate:"required"`
	Password string `json:"password" validate:"required"`
}



type CreateBookParams struct {
	Name        string `json:"name"`
	Author        string `json:"author"`
	Pages	string 	`json:"pages"`
	Content string	`json:"content"`
}

type SearchBookParams struct {
	Title string `json:"title"`
	Name        string `json:"name"`
	Pages	int8 	`json:"pages"`
}