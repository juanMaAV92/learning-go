// https://refactoring.guru/es/design-patterns/adapter

package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (c *CashPayment) Pay() {
	fmt.Println("Paying with cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}

func (b *BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying with bank account %d\n", bankAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (b *BankPaymentAdapter) Pay() {
	b.BankPayment.Pay(b.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	bank := &BankPaymentAdapter{
		BankPayment: &BankPayment{},
		bankAccount: 123456789,
	}
	ProcessPayment(bank)
}
