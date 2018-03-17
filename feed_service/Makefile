build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/rcholic/rally_bargain/feed_service \
		proto/feed/feed.proto
	docker build -t feed_service .