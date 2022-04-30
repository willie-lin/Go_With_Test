package wallet

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}
type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {

	w.balance += amount

}

func (w *Wallet) Balance() Bitcoin {
	fmt.Println("address of balance in Deposit is", &w.balance)
	return w.balance
}

// 在 Go 中，错误是值，因此我们可以将其重构为一个变量，并为其提供一个单一的事实来源。

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil

}
