FROM ubuntu:latest
LABEL authors="fran"
ENTRYPOINT ["top", "-b"]

FROM golang:1.23-alpine
WORKDIR /app/auth_service

COPY go.mod go.sum ./
RUN go mod download
COPY . .

COPY Roomko_auth/Scripts/CreateHTTPS_Key.sh ./CreateHTTPS_Key.sh
RUN chmod +x ./CreateHTTPS_Key.sh && ./CreateHTTPS_Key.sh

COPY wwwroot ./build
RUN go build -o roomko_auth ./build

EXPOSE 8832

CMD ["./roomko_auth"]