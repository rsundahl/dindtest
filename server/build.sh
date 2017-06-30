#! /bin/sh
GOOS=linux go build .
docker build -t rsundahl/server .
docker push rsundahl/server
