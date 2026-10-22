package main

import "fmt"

type Salesperson struct {
	Name         string
	Transactions int
	Sales        float64
}

func main() {

	var name string
	var transactions int
	var sales float64

	salespeople := []Salesperson{}

	totalsales := 0.0

	for i := 0; i < 3; i++ {

		fmt.Print("Enter name : ")
		fmt.Scanln(&name)

		fmt.Print("Enter number of transactions :")
		fmt.Scanln(&transactions)

		fmt.Print("Enter total sales : ")
		fmt.Scanln(&sales)

		person := Salesperson{
			Name:         name,
			Transactions: transactions,
			Sales:        sales,
		} // person is a new variable that has the collected data combined

		fmt.Print(person)

		salespeople = append(salespeople, person) //here we are adding the combined data stored in person to the slice/library salespeople
	}

	fmt.Println(salespeople)

	for _, person := range salespeople {
		fmt.Println(person.Name, person.Sales)
	} // this loop shows the name and sales of each salesperson in the salespeople slice

	for _, person := range salespeople {
		totalsales += person.Sales
	} // this loop sums up the total sales of all salespeople in the salespeople slice

	fmt.Printf("Total sales: %.2f\n", totalsales)

}
