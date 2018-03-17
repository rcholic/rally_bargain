build:
	# protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/rcholic/rally_bargain/feed_service \
		proto/feed/feed.proto
	protoc --proto_path=$(GOPATH)/src:. --micro_out=. --go_out=. ./proto/feed/*.proto
	docker build -t feed_service .