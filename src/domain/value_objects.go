package domain

type PersonData struct {
	Name string 
	Age	int		 
	Gender	*string 
	Address PersonDataAddress 
	Dependents []string  
}

type PersonDataAddress struct {
	Street string 
	Number int
	Country string 
	PostalCode int
}