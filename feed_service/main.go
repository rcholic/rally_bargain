package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/rcholic/rally_bargain/feed_service/proto/feed"
)

func main() {
	fmt.Println("hello world!")
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v\n", err)
	}

	db.AutoMigrate(&pb.Feed{})

	repo := &FeedRepository{db}

	srv := micro.NewService(
		micro.Name("rallybargain.srv.feed"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterFeedServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Fatalf("err: %v\n", err)
	}
}
