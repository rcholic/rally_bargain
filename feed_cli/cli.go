package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/micro/cmd"
	pb "github.com/rcholic/rally_bargain/feed_service/proto/feed"
	"github.com/sirupsen/logrus"
)

func main() {
	cmd.Init()

	client := pb.NewFeedServiceClient("rallybargain.srv.feed", microclient.DefaultClient)
	title := fmt.Sprintf("Title title created %v", time.Now())
	isLocal := false
	location := "http://www.163.com"
	datetime := time.Now().UnixNano() / 1000000 // milliseconds

	logrus.Printf("title = %s, isLocal = %v, location = %s, datetime = %v", title, isLocal, location, datetime)

	r, err := client.Create(context.Background(), &pb.Feed{
		Title:    title,
		IsLocal:  isLocal,
		Location: location,
		Datetime: datetime,
	})

	if err != nil {
		logrus.Fatalf("could not create: %v", err)
	}
	logrus.Infof("Created: feed id = %s", r.Feed.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		logrus.Fatalf("could not list users: %v", err)
	}

	for _, v := range getAll.Feeds {
		logrus.Println(v)
	}

	// os.Exit(0) // exit here
}
