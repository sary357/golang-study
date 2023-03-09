package payroll

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

func (m Manager) Pay() (string, float64) {
	return m.Individual.FullName(), m.Calculate()
}

func (m Manager) Calculate() float64 {
	return m.Salary + (m.Salary * m.CommissionRate)
}
