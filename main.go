package main

import (
	"fmt"
	"practice/Models"
	"practice/Mydict"
)

func main() {
	Bonus := 1000
	papers := []string{"NID","Passport"}
	moon := models.Person{Name: "moon",Age: 25,Credentials: papers}
	//fmt.Println(moon)
	newAccount := models.Accounts{AccId:"1", Owner: moon,Balance: 0}
	fmt.Println("Account openning bonus: 1000 TK added")
	newAccount.AddMoney(Bonus)
	fmt.Println("New Balance : ")
	newAccount.TotalBalance()
	fmt.Println("\nAccount Deatails ")
	newAccount.AccountInfo()
	
	Infodictionary := mydict.Dictionary{newAccount.AccId:moon.String()}

	val, err := Infodictionary.Search("1")
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(val)
	}
}