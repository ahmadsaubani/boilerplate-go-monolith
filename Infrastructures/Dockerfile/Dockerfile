FROM golang:1.25.1-alpine AS builder

LABEL maintainer="Boilerplate Dev <myname@ahmadsaubani.com>"

ENV GO111MODULE=on

WORKDIR /go/src/boilerplate-go
ADD . /go/src/boilerplate-go

RUN apk update && apk add --no-cache git ca-certificates tzdata
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o engine main.go

# Distribution
FROM alpine:3.20

RUN apk update && apk add --no-cache ca-certificates tzdata

ARG APP_ENV=Docker
ENV env_state=$APP_ENV

EXPOSE 3000

COPY --from=builder /go/src/boilerplate-go/engine /app/
COPY --from=builder /go/src/boilerplate-go/Environment /go/src/boilerplate-go/Environment

CMD /app/engine --env=$env_state