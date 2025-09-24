FROM golang:alpine

WORKDIR /build
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY . .

RUN go build -o wallet ./cmd/main.go 

EXPOSE 8080

CMD ["./wallet"]