#! /bin/sh
docker-compose down --remove-orphans -v
docker service rm `docker service ls -q` 2>/dev/null
docker secret rm `docker secret ls -q` 2>/dev/null
echo "Hit [CTRL-C] to stop docker-compose up"
sleep 5
docker-compose up --build --remove-orphans
