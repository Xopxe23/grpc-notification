package main

import (
	"context"
	"log"

	"github.com/xopxe23/grpc-notification/proto/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := notification.NewNotificationServiceClient(conn)

	response, err := c.Notify(context.Background(), &notification.NotificationRequest{
		Message: "Я на пути становления Golang Ninja",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Message: ", response.Status)
}
