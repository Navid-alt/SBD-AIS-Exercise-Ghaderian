package client

import (
	"context"
	"exc8/pb"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClient struct {
	client pb.OrderServiceClient
}

func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.NewClient(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Run() error {
	// todo
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// 1. List drinks
	fmt.Println("Requesting drinks 🍹🍺☕")
	menu, err := c.client.GetDrinks(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("failed to get drinks: %w", err)
	}

	fmt.Println("Available drinks:")
	for _, d := range menu.Drinks {
		// Logic: If price is 0, don't print the price tag
		if d.Price > 0 {
			fmt.Printf("   > id:%d  name:\"%s\"  price:%.0f  description:\"%s\"\n", d.Id, d.Name, d.Price, d.Description)
		} else {
			// Coffee matches this (Price is 0 or undefined), so we skip the price part
			fmt.Printf("   > id:%d  name:\"%s\"  description:\"%s\"\n", d.Id, d.Name, d.Description)
		}
	}
	orderItem := func(id int32, qty int32, name string) {
		fmt.Printf("   > Ordering: %d x %s\n", qty, name)
		// We ignore the boolean response here as per the simple output requirement,
		// but in a real app you'd check if success.Value is true.
		_, err := c.client.OrderDrink(ctx, &pb.OrderRequest{
			DrinkId:  id,
			Quantity: qty,
		})
		if err != nil {
			fmt.Printf("Error ordering %s: %v\n", name, err)
		}
	}
	// 2. Order a few drinks
	fmt.Println("Ordering drinks 👨‍🍳⏱️🍻🍻")
	orderItem(1, 2, "Spritzer")
	orderItem(2, 2, "Beer")
	orderItem(3, 2, "Coffee")
	// 3. Order more drinks
	fmt.Println("Ordering another round of drinks 👨‍🍳⏱️🍻🍻")
	orderItem(1, 6, "Spritzer")
	orderItem(2, 6, "Beer")
	orderItem(3, 6, "Coffee")
	// 4. Get order total
	//
	fmt.Println("Getting the bill 💹💹💹")
	bill, err := c.client.GetOrders(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("failed to get bill: %w", err)
	}

	for _, item := range bill.Orders {
		fmt.Printf("   > Total: %d x %s\n", item.Quantity, item.Drink.Name)
	}
	// print responses after each call
	return nil
}
