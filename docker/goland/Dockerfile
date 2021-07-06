FROM golang:latest

# ----- SETUP -----

# Enable Go modules
ENV GO111MODULE=on

# Set the current working with go absolute path
WORKDIR /go/src/github.com/GoElasticsearch/api

# Install 'air' live-reload go module
RUN go get -u github.com/cosmtrek/air

# Use the excutable
ENTRYPOINT ["/go/bin/air"]