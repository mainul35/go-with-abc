package handling_json

type Employee struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	Profession string `json:"profession"`
}
