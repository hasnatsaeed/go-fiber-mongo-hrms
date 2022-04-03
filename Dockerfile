#  base image for Go
FROM golang:latest
# Set the Current Working Directory inside the container

WORKDIR /app

RUN git clone https://github.com/hasnatsaeed/go-fiber-mongo-hrms.git

WORKDIR /app/go-fiber-mongo-hrms

EXPOSE 9010

ENTRYPOINT ["go", "run", "/app/go-fiber-mongo-hrms/cmd/main/main.go"]