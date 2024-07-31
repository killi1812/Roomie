FROM ubuntu:latest
LABEL authors="fran"
ENTRYPOINT ["top", "-b"]

FROM golang:1.20-alpine
WORKDIR /app/auth_service

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o roomko_auth ./

EXPOSE 8832

CMD ["./roomko_auth"]