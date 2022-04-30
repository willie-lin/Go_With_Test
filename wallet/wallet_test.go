package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	}

	assertError := func(t *testing.T, got error, want string) {
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))
		if err != nil {
			return
		}

		assertBalance(t, wallet, Bitcoin(10))
		assertError(t, err, "cannot withdraw, insufficient funds")
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)

		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
		assertError(t, err, "cannot withdraw, insufficient funds")
	})

}
