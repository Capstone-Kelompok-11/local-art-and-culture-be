package request

type Auth struct {
	Name 		string 	`json:"name"`
	Password	string	`json:"password"`
}