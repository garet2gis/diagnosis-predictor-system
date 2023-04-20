# Step 1: Modules caching
FROM golang:1.19-alpine as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.19-alpine as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /api-gateway
WORKDIR /api-gateway
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/main ./cmd/app/

# Step 3: Final
FROM alpine:3.17

COPY --from=builder /api-gateway/.bin/main .

ENTRYPOINT ["./main"]