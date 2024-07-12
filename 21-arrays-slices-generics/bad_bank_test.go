package main

import "testing"

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Tiago",
			To:   "Suian",
			Sum:  100,
		},
		{
			From: "Éber",
			To:   "Tiago",
			Sum:  25,
		},
	}

	AssertEqual(t, BalanceFor(transactions, "Suian"), 100)
	AssertEqual(t, BalanceFor(transactions, "Tiago"), -75)
	AssertEqual(t, BalanceFor(transactions, "Éber"), -25)
}
