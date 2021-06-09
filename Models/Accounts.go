package models

import "fmt"

type Accounts struct {
	AccId string
	Owner   Person
	Balance int
}

func (a *Accounts) AddMoney(amount int) {
	a.Balance += amount
}

func (a Accounts) TotalBalance() {
	fmt.Println(a.Balance)
}

func (a Accounts) AccountInfo(){
	fmt.Println("Name : ",a.Owner.Name,
	"\nAge : ", a.Owner.Age,
	"\nBalance: ",a.Balance)
}


