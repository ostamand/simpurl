SHELL := /bin/bash
test:
	go test ./... -count=1
setup:
	gcloud auth configure-docker $(REGION)-docker.pkg.dev
run:
	go build -o server ./web && ./server
docker:
	docker build --build-arg config_file=$(CONFIG_FILE) -t $(TAG) .
docker-run:
	docker run --rm -p $(PORT):$(PORT) -p $(DB_PORT):$(DB_PORT) $(TAG)
docker-push:
	docker push $(TAG)
deploy: docker docker-push
	gcloud run deploy shorturl --image $(TAG) --max-instances $(INSTANCES) --region=$(REGION)