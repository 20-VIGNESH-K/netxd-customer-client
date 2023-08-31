package main

import (
	"context"
	"fmt"
	"log"

	cus "github.com/20-VIGNESH-K/netxd-customer-proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	client := cus.NewCustomerServiceClient(conn)

	response, err := client.CreateCustomer(context.Background(), &cus.Customer{
		CustomerId: 187,
		FirstName:  "sabari",
		LastName:   "S",
		BankId:     79,
		Balance:    7540.0,
		CreatedAt:  "",
		UpdatedAt:  "",
		IsActive:   true,
	})
	if err != nil {
		log.Fatalf("failed to call CreateCustomer: %v", err)
	}
	fmt.Printf("Response: %s\n", response)
}
