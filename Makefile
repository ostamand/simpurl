SHELL := /bin/bash
down: 
	docker compose down
up: down
	docker compose up -d --build
test:
	go test ./... -count=1
setup:
	gcloud auth configure-docker $(REGION)-docker.pkg.dev
build:
	go build -o server ./cmd/web
run: build
	./server
docker:
	docker build -f Dockerfile-server --build-arg config_file=$(CONFIG_FILE) -t $(TAG) .
docker-push:
	docker push $(TAG)
deploy: docker docker-push
	gcloud run deploy shorturl --image $(TAG) --max-instances $(INSTANCES) --region=$(REGION)