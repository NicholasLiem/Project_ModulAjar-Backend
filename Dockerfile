#Pulling Golang Image
FROM golang:1.20-alpine as BuildStage

#Setting workdir in /app
WORKDIR /app

#Copying go mod and sum files
COPY go.mod go.sum ./

#Installing modules
RUN go mod download

#Copying source code and env
COPY . ./

#Compiling app
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend-app


# Deploy Stage

FROM alpine:latest

WORKDIR /

COPY --from=BuildStage /backend-app /backend-app

COPY .env .

#Setting up the port
EXPOSE 8080

RUN addgroup -g 1000 nonroot && adduser -u 1000 -G nonroot -s /bin/sh -D nonroot

USER nonroot

#Running the app
CMD ["/backend-app"]