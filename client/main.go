package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	pb "github.com/tattoo1880/my-grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	//conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer conn.Close()

	newClient, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer newClient.Close()

	client := pb.NewUserServiceClient(newClient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetUser(ctx, &pb.UserRequest{Id: 100})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("User:", resp.Name)
}
