package datamodel

type Employee struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	CompanyName string `json:"companyName"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Postal      string `json:"postal"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Web         string `json:"web"`
}
