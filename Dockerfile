# from: https://github.com/GoogleCloudPlatform/golang-samples/blob/main/run/helloworld/Dockerfile
FROM golang:1.17.5-buster

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

ARG config_file 
ENV CONFIG_FILE=${config_file}
COPY ./${config_file} .

RUN go build -o server ./cmd/web
CMD ["/app/server"]