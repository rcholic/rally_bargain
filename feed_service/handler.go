package main

import (
	"golang.org/x/net/context"

	pb "github.com/rcholic/rally_bargain/feed_service/proto/feed"
)

type service struct {
	repo Repository
}

func (srv *service) Create(ctx context.Context, req *pb.Feed, res *pb.Response) error {
	if err := srv.repo.Create(req); err != nil {
		return err
	}

	res.Feed = req

	return nil
}

func (srv *service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	feeds, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Feeds = feeds
	return nil
}
