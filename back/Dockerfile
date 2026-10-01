FROM golang:alpine as builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app

FROM scratch

COPY --from=builder /app /app
ENTRYPOINT ["/app"]