package main

import "testing"

func TestBadBank(t *testing.T) {
	var (
		suian = Account{Name: "Suian", Balance: 100}
		tiago = Account{Name: "Tiago", Balance: 75}
		eber  = Account{Name: "Éber", Balance: 200}

		transactions = []Transaction{
			NewTransaction(tiago, suian, 100),
			NewTransaction(eber, tiago, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(suian), 200)
	AssertEqual(t, newBalanceFor(tiago), 0)
	AssertEqual(t, newBalanceFor(eber), 175)
}
