#!/bin/bash

output_dir=${1:-.}


# Check if the output directory exists
if [ ! -d "$output_dir" ]; then
  mkdir -p "$output_dir"
  echo "Created directory: $output_dir"
fi

# Check if the keys directory exists
keys_dir="$output_dir/keys"
if [ ! -d "$keys_dir" ]; then
  mkdir -p "$keys_dir"
  echo "Created directory: $keys_dir"
fi

# TODO create them even though they exist
if [ ! -f "$keys_dir/Https_key.pem" ]; then
  echo "Creating HTTPS key"
  openssl req -x509 -newkey rsa:4096 -keyout "$keys_dir/Https_key.pem" -out "$keys_dir/Https_cert.pem" -days 365 -subj "/C=HR/ST=Zagreb/L=Zagreb/O=Roomko/CN=Roomko.com" -nodes
fi