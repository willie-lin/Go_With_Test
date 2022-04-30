package wallet

import "fmt"

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(amount int) {

	w.balance += amount

}

func (w Wallet) Balance() int {
	fmt.Println("address of balance in Deposit is", &w.balance)
	return w.balance

}
