package main

import (
	"context"
	"log"
	"time"

	"github.com/hanhuquynh/grpc/pb"
	"github.com/rs/xid"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:2001", grpc.WithInsecure())

	if err != nil {
		log.Fatal("err: ", err)
	}

	defer cc.Close()

	log.Println("Client...")

	client := pb.NewUserServiceClient(cc)

	insert(client)
	// read(client)
	// update(client)
	// delete(client)
}

func insert(client pb.UserServiceClient) {
	req := &pb.InsertRequest{
		User: &pb.UserPartner{
			Id:          xid.New().String(),
			UserId:      "4",
			PartnerId:   "4",
			AliasUserId: "4",
			Phone:       "0988776654",
			Created:     time.Now().UnixMilli(),
			UpdatedAt:   time.Now().UnixMilli(),
		},
	}
	resp, err := client.Insert(context.Background(), req)

	if err != nil {
		log.Printf("insert err %v: ", err)
		return
	}

	log.Printf("Insert response %v: ", resp)
}

func read(client pb.UserServiceClient) {
	resp, err := client.Read(context.Background(), &pb.ReadRequest{})
	if err != nil {
		log.Printf("call read err %v\n", err)
		return
	}

	log.Printf("read response %+v\n", resp)
}

func update(client pb.UserServiceClient) {
	resp, err := client.Update(context.Background(), &pb.UpdateRequest{
		NewUser: &pb.UserPartner{
			UserId:    "2",
			PartnerId: "2",
			Phone:     "0900986905",
		},
	})

	if err != nil {
		log.Fatalf("update user %v: %v", resp, err)
	}

	log.Println("Update: ", resp)
}

func delete(client pb.UserServiceClient) {
	resp, err := client.Delete(context.Background(), &pb.DeleteRequest{
		User: &pb.UserPartner{
			UserId: "4",
		},
	})

	if err != nil {
		log.Fatalf("delete user %v: %v", resp, err)
	}

	log.Println("Delete", resp)
}
