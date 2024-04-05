FROM golang:latest

WORKDIR /app

# COPY go.mod .
# COPY go.sum .
# RUN go mod download
# RUN go build -trimpath -ldflags="-s -w" -o app .

RUN go install github.com/cosmtrek/air@latest
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest

# CMD ["app"]
