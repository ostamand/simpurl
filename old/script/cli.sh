docker compose up -d --build
docker compose exec server ping -c 10 db
docker compose exec server go test ./... -count=1
docker compose down