//go get github.com/shopspring/decimal
package tests

import (
	"../models"
	"../modules/billing"
	"../modules/storage"
	"fmt"
	"time"
)


func main() {
	accountStorage := storage.NewInMemoryDB()

	account1 := models.Account{}
	account1.Init(1, "Mark")
	accountStorage.Create(account1.GetId(), &account1)
	account1.SetAmount("2454354358735793579823794723875982738472387423759872385792374283748723947238749237.33")

	fmt.Println(account1.GetId())
	fmt.Println(account1.GetName())
	fmt.Println(account1.GetAmount().String())

	account1.GetAmount().Add("22.02")
	fmt.Println(account1.GetAmount().String())

	account1.GetAmount().Sub("53.07")
	fmt.Println(account1.GetAmount().String())

	account2 := models.Account{}
	account2.Init(2, "Archi")
	accountStorage.Create(account2.GetId(), &account2)
	account2.SetAmount("1000")

	fmt.Println("============before=================")
	fmt.Println(account1.GetAmount().String())
	fmt.Println(account2.GetAmount().String())

	for i := 0; i <= 100000; i++ {
		go billing.PayToAccount(&account1, &account2, "1.01")
	}

	//billing.PayToAccount(&account1, &account2, "1.01")
	time.Sleep(time.Second)

	fmt.Println("============after==================")
	fmt.Println(account1.GetAmount().String())
	fmt.Println(account2.GetAmount().String())
}