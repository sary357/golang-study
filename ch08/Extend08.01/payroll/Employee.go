package payroll

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}
