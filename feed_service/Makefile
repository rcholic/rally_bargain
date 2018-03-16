build:
	protoc -I. --go_out=plugins=micro:/Users/guoliangwang/Desktop/Code/Go/Workspace/src/github.com/rcholic/rally_bargain/feed_service \
		proto/feed/feed.proto
	docker build -t feed_service .