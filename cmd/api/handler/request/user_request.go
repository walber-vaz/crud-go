package request

type UserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=60,containsany=!@#$%^&*()_+"`
	FirstName string `json:"first_name" binding:"required,min=2,max=60"`
	LastName  string `json:"last_name" binding:"required,min=2,max=60"`
	Age       int8   `json:"age" binding:"required,numeric,min=18,max=140"`
}
