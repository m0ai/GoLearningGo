package main

import (
	"fmt"
	"github.com/m0ai/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("moai")
	account.Despoit(10)
	fmt.Println(account.Balance())

	account.ChangeOwner("nico")

	fmt.Println(account)
}
