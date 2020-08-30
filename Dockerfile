FROM golang:1.14 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN GOOS=linux CGO_ENABLED=0 go build ./cmd/store

FROM alpine
WORKDIR /app
COPY --from=builder /app/store /app/assets ./
CMD [ "/app/store", "serve" ]