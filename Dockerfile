FROM golang:latest

COPY . /ai-kcal-agent
WORKDIR /ai-kcal-agent

RUN go mod download
RUN go build -o bin/ cmd/main.go
EXPOSE 8081
CMD ["/ai-kcal-agent/bin/main"]