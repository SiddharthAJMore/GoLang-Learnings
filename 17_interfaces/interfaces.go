package main

import "fmt"

type payment1 struct {
	gateway razorPay1
}

func (p payment1) makePayment(amount float32) {
	//razorPay1PaymentGW := razorPay{}
	//razorPay1PaymentGW.pay(amount)

	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)

	p.gateway.pay(amount)
}

type razorPay1 struct{}

func (r razorPay1) pay(amount float32) {
	fmt.Println("Making payment1 using RazorPay", amount)
}

type stripe1 struct{}

func (s stripe1) pay(amount float32) {
	fmt.Println("Making payment1 using Stripe", amount)
}

func main() {
	//stripePaymentGw := stripe1{}
	razorPay1PaymentGw := razorPay1{}

	payMoney := payment1{
		razorPay1PaymentGw,
	}
	payMoney.makePayment(100)
}
