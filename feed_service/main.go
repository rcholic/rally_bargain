package main

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro"
	pb "github.com/rcholic/rally_bargain/feed_service/proto/feed"
)

func main() {
	fmt.Println("hello world!")
	db, err := CreateConnection()
	defer db.Close()
	if err != nil {
		fmt.Printf("Could not connect to DB: %v", err)
		// log.Fatalf("Could not connect to DB: %v\n", err)
	}

	db.AutoMigrate(&pb.Feed{})

	repo := &FeedRepository{db}

	srv := micro.NewService(
		micro.Name("rallybargain.srv.feed"),
		micro.Version("latest"),
		micro.RegisterInterval(time.Second*10),
	)

	srv.Init()

	pb.RegisterFeedServiceHandler(srv.Server(), &service{repo})

	log.Println("service server is running...")
	if err := srv.Run(); err != nil {
		log.Fatalf("err: %v\n", err)
	}
}
