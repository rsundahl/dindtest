#! /bin/sh
GOOS=linux go build .
docker build -t rsundahl/api .
docker push rsundahl/api
