package response

type Address struct {
	Id			uint	`json:"id"`
	Address 	string	`json:"address"`
	Location	string	`json:"location"`
	Users   []Users   	
	Creator Creator 
}