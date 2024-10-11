FROM ubuntu:latest
LABEL authors="fran"
ENTRYPOINT ["top", "-b"]

FROM golang:1.23-alpine
WORKDIR /app/auth_service

RUN mkdir "src"

COPY go.mod go.sum ./src/
WORKDIR ./src
RUN go mod download
COPY . .
RUN go build -o ../build/roomko_auth 
WORKDIR ../build

# Copy the script into the Docker image
COPY Scripts/CreateHTTPS_Key.sh ./CreateHTTPS_Key.sh
RUN apk add --no-cache openssl
# Make the script executable and run it
RUN chmod +x ./CreateHTTPS_Key.sh
RUN ["/bin/sh", "./CreateHTTPS_Key.sh"]
RUN cp -R ../src/wwwroot ./

EXPOSE 8832
CMD ["./roomko_auth"]