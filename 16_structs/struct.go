package main

import (
	"fmt"
	"time"
)

// customer struct
type customer struct {
	name   string
	mobile string
}

// order struct
type order struct {
	id       string
	amount   float32
	status   string
	createAt time.Time
	customer
}

// functions of the class
func (o *order) setStatus(status string) {
	o.status = status // dereferencing handled by struct automatically
}

func (o *order) getAmount() float32 {
	return o.amount
}

// constructor
func newOrder(id string, amount float32, status string, createdAt time.Time) order {
	return order{
		id:       id,
		amount:   amount,
		status:   status,
		createAt: createdAt,
	}
}

// structs for 1 time usage -- given in below main function

func main() {

	myOrder1 := order{
		id:     "1",
		amount: 100.00,
		status: "Delivered",
	}

	myOrder1.createAt = time.Now()

	fmt.Println("My order 1", myOrder1)

	myOrder2 := order{
		id:     "2",
		amount: 50.00,
		status: "Pending",
	}

	//myOrder2.status = "Shipped"
	myOrder2.setStatus("Shipped")

	fmt.Println("Order 2 amount", myOrder2.amount)
	fmt.Println("My order 2", myOrder2)

	myOrder3 := newOrder("3", 30.50, "Pending", time.Now())
	fmt.Println("My order 3", myOrder3)

	//struct for 1 time usage
	someStruct := struct {
		name    string
		message string
	}{
		name:    "Siddharth",
		message: "Hello world",
	}
	fmt.Println("One time struct", someStruct)

	myOrder4 := newOrder("4", 30.50, "Pending", time.Now())
	myOrder4.customer.name = "Siddharth"
	myOrder4.name = "Shraddha" // automatically sets on customer

	fmt.Println("My order 4", myOrder4)
	fmt.Println("My order 4", myOrder4.customer.name)

}
