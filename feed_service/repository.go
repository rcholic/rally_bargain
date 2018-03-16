package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/rcholic/rally_bargain/feed_service/proto/feed"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Create(feed *pb.Feed) error
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
