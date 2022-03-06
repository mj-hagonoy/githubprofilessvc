# Pull base image
FROM golang:1.16-alpine as builder
WORKDIR /github.com/mj-hagonoy/githubprofilessvc
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/main ./cmd
RUN ls 

FROM alpine:latest
ENV GOFLAGS=-mod=vendor
WORKDIR /github.com/mj-hagonoy/githubprofilessvc
COPY config.yaml ./
COPY --from=builder /github.com/mj-hagonoy/githubprofilessvc/build/main .

EXPOSE 8080
CMD ["./main","--config", "./config.yaml"]