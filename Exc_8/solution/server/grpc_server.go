package server

import (
	"context"
	"exc8/pb"
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer
	drinks map[int32]*pb.Drink
	orders map[int32]int32
	mu     sync.Mutex
}

func StartGrpcServer() error {
	// Create a new gRPC server.
	srv := grpc.NewServer()
	// Create grpc service
	grpcService := &GRPCService{
		drinks: make(map[int32]*pb.Drink),
		orders: make(map[int32]int32),
	}

	grpcService.drinks[1] = &pb.Drink{Id: 1, Name: "Spritzer", Price: 2, Description: "Wine with soda"}
	grpcService.drinks[2] = &pb.Drink{Id: 2, Name: "Beer", Price: 3, Description: "Hagenberger Gold"}
	grpcService.drinks[3] = &pb.Drink{Id: 3, Name: "Coffee", Description: "Mifare isn't that secure"}
	// Register our service implementation with the gRPC server.
	pb.RegisterOrderServiceServer(srv, grpcService)
	// Serve gRPC server on port 4000.
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}
	err = srv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

// todo implement functions
func (s *GRPCService) GetDrinks(ctx context.Context, _ *emptypb.Empty) (*pb.DrinkList, error) {
	var list []*pb.Drink
	// Loop 1..3 to maintain order
	for i := int32(1); i <= 3; i++ {
		if d, ok := s.drinks[i]; ok {
			list = append(list, d)
		}
	}
	return &pb.DrinkList{Drinks: list}, nil
}

func (s *GRPCService) OrderDrink(ctx context.Context, req *pb.OrderRequest) (*wrapperspb.BoolValue, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if drink exists
	if _, exists := s.drinks[req.DrinkId]; !exists {
		return &wrapperspb.BoolValue{Value: false}, fmt.Errorf("drink with id %d not found", req.DrinkId)
	}

	// Add quantity to orders
	s.orders[req.DrinkId] += req.Quantity
	return &wrapperspb.BoolValue{Value: true}, nil
}

func (s *GRPCService) GetOrders(ctx context.Context, _ *emptypb.Empty) (*pb.OrderList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var orderItems []*pb.OrderItem
	// Loop 1..3 to maintain order
	for i := int32(1); i <= 3; i++ {
		qty, ordered := s.orders[i]
		if ordered && qty > 0 {
			orderItems = append(orderItems, &pb.OrderItem{
				Drink:    s.drinks[i],
				Quantity: qty,
			})
		}
	}
	return &pb.OrderList{Orders: orderItems}, nil
}
