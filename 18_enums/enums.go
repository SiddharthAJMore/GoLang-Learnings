package main

import "fmt"

// enumerated types

type OrderStatus int

const (
	RECEIVED OrderStatus = iota
	CONFIRMED
	PREPARED
	DISPATCHED
	DELIVERED
)

type Role string

const (
	ADMIN    Role = "admin"
	OPERATOR      = "operator"
	USER          = "user"
)

func changeOrderStatusTo(newStatus OrderStatus) {
	fmt.Println("Changing order status to", newStatus)
}

func main() {

	changeOrderStatusTo(RECEIVED)

	fmt.Println("admin role is", ADMIN)
}
