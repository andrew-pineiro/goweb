# syntax=docker/dockerfile:1

FROM golang:1.22

# Set destination for COPY
WORKDIR /www

# Download Go modules
RUN cd ../src
COPY go.mod go.sum ./
RUN go mod download

COPY .. ./

# BUILD
RUN CGO_ENABLED=0 GOOS=linux go build -o ../goweb

# THIS SHOULD MATCH --Port argument
EXPOSE 8080

# Run
RUN cd ..
RUN chmod +x ./goweb
CMD ["./goweb"]