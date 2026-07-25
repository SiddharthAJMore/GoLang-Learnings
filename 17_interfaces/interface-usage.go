package main

import "fmt"

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type paymenter interface {
	pay(amount float32)
}

type razorPay struct{}

func (r razorPay) pay(amount float32) {
	fmt.Println("Making payment using RazorPay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using Stripe", amount)
}
func main() {
	//razorPayPaymentGw := razorPay{}
	stripePaymentGw := stripe{}

	paymentGw := payment{
		stripePaymentGw,
	}

	paymentGw.makePayment(300)
}
