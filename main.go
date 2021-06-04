package main

import (
	"practice/Models"
	"fmt"
)

func main() {
	Bonus := 1000
	papers := []string{"NID","Passport"}
	moon := models.Person{Name: "moon",Age: 25,Credentials: papers}
	newAccount := models.Accounts{Owner: moon,Balance: 0}
	fmt.Println("Account openning bonus: 1000 TK added")
	newAccount.AddMoney(Bonus)
	fmt.Println("New Balance : ")
	newAccount.TotalBalance()
	fmt.Println("\nAccount Deatails ")
	newAccount.AccountInfo()

}