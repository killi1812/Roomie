#!/bin/bash

openssl req -x509 -nodes -days 365 -newkey rsa:4096 -keyout ./../build/keys/Https_key.pem -out ./../build/keys/Https_cert.pem
