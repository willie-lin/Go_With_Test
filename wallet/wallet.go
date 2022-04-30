package wallet

import "fmt"

type Wallet struct {
	balance Bitcoin
}
type Bitcoin int

func (w *Wallet) Deposit(amount Bitcoin) {

	w.balance += amount

}

func (w *Wallet) Balance() Bitcoin {
	fmt.Println("address of balance in Deposit is", &w.balance)
	return w.balance
}
