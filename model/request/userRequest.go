package request

type UserRequest struct {
	Name     string    `json:"name"`
	Age      int8      `json:"age"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Address  []Address `json:"address"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
	Cep     string `json:"cep"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
