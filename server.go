package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/omar-khawaja/grpc-calculator-opentracing/lib/tracing"
	pb "github.com/omar-khawaja/grpc-calculator/calculator"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s server) Multiply(ctx context.Context, numbers *pb.Numbers) (*pb.Result, error) {
	var product int32 = numbers.Operand[0]
	for i := 1; i < len(numbers.Operand); i++ {
		product *= numbers.Operand[i]
	}
	return &pb.Result{Result: product}, nil
}

func (s server) Add(ctx context.Context, numbers *pb.Numbers) (*pb.Result, error) {
	var sum int32 = numbers.Operand[0]
	for i := 1; i < len(numbers.Operand); i++ {
		sum += numbers.Operand[i]
	}
	return &pb.Result{Result: sum}, nil
}

func (s server) Divide(ctx context.Context, numbers *pb.Numbers) (*pb.Result, error) {
	var quotient int32 = numbers.Operand[0]
	for i := 1; i < len(numbers.Operand); i++ {
		if numbers.Operand[i] == 0 {
			return &pb.Result{Result: 0}, errors.New("Sorry. You cannot divide by 0")
		}
		quotient /= numbers.Operand[i]
	}
	return &pb.Result{Result: quotient}, nil
}

func (s server) Subtract(ctx context.Context, numbers *pb.Numbers) (*pb.Result, error) {
	var difference int32 = numbers.Operand[0]
	for i := 1; i < len(numbers.Operand); i++ {
		difference -= numbers.Operand[i]
	}
	return &pb.Result{Result: difference}, nil
}

func main() {
	tracer, closer := tracing.Init("calculator-trace")
	defer closer.Close()
	addr := "127.0.0.1"
	port := 8080
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		log.Println(err)
		return
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(tracer)),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer)))
	pb.RegisterCalculatorServer(s, server{})
	log.Printf("Starting server on %d\n", port)
	s.Serve(l)
	if err := s.Serve(l); err != nil {
		log.Println(err)
		return
	}
}
