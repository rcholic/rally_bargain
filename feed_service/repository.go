package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/rcholic/rally_bargain/feed_service/proto/feed"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Create(feed *pb.Feed) error
	GetAll() ([]*pb.Feed, error)
}

type FeedRepository struct {
	db *gorm.DB
}

func (repo *FeedRepository) Create(feed *pb.Feed) error {
	if err := repo.db.Create(feed).Error; err != nil {
		return err
	}
	logrus.Infof("feed created!")

	return nil
}

func (repo *FeedRepository) GetAll() ([]*pb.Feed, error) {
	var feeds []*pb.Feed
	if err := repo.db.Find(&feeds).Error; err != nil {
		return nil, err
	}

	return feeds, nil
}
