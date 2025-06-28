package main

import "fmt"

type Account struct {
	Owner   string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount float64) {
	if amount > a.Balance {
		fmt.Println("Недостаточно средств.")
		return
	}
	a.Balance -= amount
}

func (a Account) GetBalance() float64 {
	return a.Balance
}

func main() {
	acc := Account{Owner: "Иван", Balance: 1000.0}
	fmt.Println("Баланс:", acc.GetBalance())

	acc.Deposit(500)
	fmt.Println("После пополнения:", acc.GetBalance())

	acc.Withdraw(300)
	fmt.Println("После снятия:", acc.GetBalance())

	acc.Withdraw(1500)
}
