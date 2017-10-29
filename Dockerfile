FROM golang:latest

ARG CACHEBUST=1
EXPOSE 4569
COPY main.go /app/main.go
COPY static/ /app/static

CMD cd /app && go run /app/main.go
