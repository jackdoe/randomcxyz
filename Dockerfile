FROM golang:latest

ARG CACHEBUST=1
EXPOSE 4569
COPY main.go /app/main.go
COPY static/ /app/static

CMD go run /app/main.go
