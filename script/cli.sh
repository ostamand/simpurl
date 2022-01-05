docker compose up -d --build
docker compose run server ping -c 10 db
docker compose run server go test ./... -count=1
docker compose down