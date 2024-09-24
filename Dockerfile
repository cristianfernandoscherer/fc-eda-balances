FROM golang:1.19
WORKDIR /app

RUN apt-get update && apt-get install -y librdkafka-dev
COPY go.mod ./
RUN go mod download
COPY . .

RUN go build -o /cmd/app ./cmd
RUN go build -o /cmd/app2 ./cmd/api

EXPOSE 3000
CMD ["sh", "-c", "/cmd/app & /cmd/app2"]
#CMD ["/cmd/app"]