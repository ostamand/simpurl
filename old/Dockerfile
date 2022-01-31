FROM golang:1.17.5-buster

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

ARG config_file 
ENV CONFIG_FILE=${config_file}
COPY ./${config_file} .

RUN go build -o server ./cmd/api
RUN go build -o simpurl ./cmd/cli

CMD ["/app/server"]