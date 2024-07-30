#!/bin/bash
# Check if the build directory exists
if [ ! -d "./../build" ]; then
  mkdir ./build
  echo "Created directory: build"
fi

# Check if the build/keys directory exists
if [ ! -d "./../build/keys" ]; then
  mkdir ./build/keys
  echo "Created directory: build/keys"
fi

# TODO create them even though they exist
if [ ! -d "./../build/keys/Https_key.pem" ]; then
  echo "Creating HTTPS key"
  openssl req -x509 -nodes -days 365 -newkey rsa:4096 -keyout ./../build/keys/Https_key.pem -out ./../build/keys/Https_cert.pem
fi

