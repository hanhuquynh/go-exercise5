package main

import (
	"context"
	"fmt"
	"log"
	"net"

	b5 "github.com/hanhuquynh/b5"
	"github.com/hanhuquynh/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UserServiceServer
}

func (server) Insert(ctx context.Context, req *pb.InsertRequest) (*pb.InsertResponse, error) {
	log.Println("Call insert user...")
	u := b5.ConvertPbUser(req.User)

	err := u.Insert()

	if err != nil {
		resp := &pb.InsertResponse{
			StatusCode: -1,
			Message:    fmt.Sprintf("Insert %+v  %v: ", u, err),
		}

		return resp, nil
	}
	resp := &pb.InsertResponse{
		StatusCode: 1,
		Message:    "OK",
	}
	return resp, nil
}

func (server) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	log.Println("Call read user...")
	data, err := b5.Read()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Read user err: %v", err)
	}

	listUser := []*pb.UserPartner{}

	for _, u := range data {
		pbUser := b5.ConvertUserPb(u)
		listUser = append(listUser, pbUser)
	}
	return &pb.ReadResponse{
		User: listUser,
	}, nil
}

func (server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	log.Println("Call update user...")
	u := b5.ConvertPbUser(req.NewUser)
	err := u.Update()
	if err != nil {
		log.Fatal(err)
	}

	resp := &pb.UpdateResponse{
		StatusCode: 1,
		Message:    "Update successfully",
	}
	return resp, nil
}

func (server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	log.Println("Call delete user...")
	u := b5.ConvertPbUser(req.User)

	err := u.Delete()
	if err != nil {
		log.Fatal(err)
	}

	resp := &pb.DeleteResponse{
		StatusCode: 1,
		Message:    "Delete Successfully",
	}
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:2001")

	if err != nil {
		log.Fatal("Err:", err)
	}

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &server{})

	log.Println("Server is running...")

	err = s.Serve(lis)

	if err != nil {
		log.Fatal("Err:", err)
	}
}
