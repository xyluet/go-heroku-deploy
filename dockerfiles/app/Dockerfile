FROM golang:1.16-alpine as builder
RUN go version
RUN apk update && apk add --no-cache git
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o=bin/app ./cmd/app/*.go

FROM alpine
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /src/bin /bin
COPY --from=builder /src/dockerfiles/app/start.sh start.sh
USER nobody:nobody
ENTRYPOINT [ "/bin/app" ]
