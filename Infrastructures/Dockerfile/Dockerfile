# Stage build
FROM golang:1.25.1-alpine AS builder

LABEL maintainer="Boilerplate Dev <myname@ahmadsaubani.com>"

ENV GO111MODULE=on

WORKDIR /go/src/boilerplate-go

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apk add --no-cache git ca-certificates tzdata

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -trimpath \
    -ldflags="-s -w" \
    -o api-go main.go

FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata wget

ARG APP_ENV=production
ENV env_state=$APP_ENV

EXPOSE 3000

COPY --from=builder /go/src/boilerplate-go/api-go /api-go
COPY --from=builder /go/src/boilerplate-go/Environment /go/src/boilerplate-go/Environment

WORKDIR /

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:3000/health || exit 1

ENTRYPOINT ["/api-go"]
CMD ["--env=production"]