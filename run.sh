#!/usr/bin/bash
export DATABASE_URL="host=127.0.0.1 user=postgres password=postgres dbname=yenepay-demo sslmode=disable"
docker kill $(docker ps -q --all)
docker container rm $(docker container ls -q --all)
docker run --name postgresql-container -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=yenepay-demo -d postgres
rm -rf target
revel run -a .
