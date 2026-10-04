package main

import "fmt"

type Employee struct {
	Name        string
	Salary      float64
	YearsWorked int
}

func (e Employee) Bonus() float64 {
	periods := e.YearsWorked / 5 // целочисленное деление (например, 13 / 5 = 2)
	return e.Salary * (float64(periods) * 0.10)
}

type Company []Employee

func (c Company) TotalPayroll() float64 {
	var total float64
	for _, emp := range c {
		total += emp.Salary + emp.Bonus()
	}
	return total
}

func (c Company) TopEarner() Employee {
	if len(c) == 0 {
		return Employee{}
	}

	top := c[0]
	topTotal := top.Salary + top.Bonus()

	for _, emp := range c {
		currentTotal := emp.Salary + emp.Bonus()
		if currentTotal > topTotal {
			top = emp
			topTotal = currentTotal
		}
	}

	return top
}

func main() {
	company := Company{
		{Name: "Али", Salary: 1000, YearsWorked: 5},
		{Name: "Умед", Salary: 2500, YearsWorked: 10},
		{Name: "Бахтиёр", Salary: 1500, YearsWorked: 13},
		{Name: "Аминчон", Salary: 500, YearsWorked: 4},
		{Name: "Улугбек", Salary: 5000, YearsWorked: 20},
	}

	fmt.Println("Информация о сотрудниках:")
	for _, emp := range company {
		bonus := emp.Bonus()
		total := emp.Salary + bonus
		fmt.Printf("- %s: Оклад = %.2f, Бонус (%d лет) = %.2f, Итого = %.2f\n",
			emp.Name, emp.Salary, emp.YearsWorked, bonus, total)
	}

	fmt.Printf("\nОбщие расходы на зарплаты (Total Payroll): %.2f\n", company.TotalPayroll())

	top := company.TopEarner()
	fmt.Printf("Самый высокооплачиваемый сотрудник: %s (Всего: %.2f с учётом бонуса)\n",
		top.Name, top.Salary+top.Bonus())
}
