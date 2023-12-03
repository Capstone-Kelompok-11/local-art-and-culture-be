package request

type Address struct {
	Id			uint	`json:"id"`
	Address 	string	`json:"address"`
	Location	string	`json:"location"`
}