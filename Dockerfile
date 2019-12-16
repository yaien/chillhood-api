FROM golang:1.12 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN GOOS=linux CGO_ENABLED=0 go build ./cmd/server

FROM scratch
WORKDIR /app
COPY --from=builder /app/server .
ENTRYPOINT [ "/app/server" ]