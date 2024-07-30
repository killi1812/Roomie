#!/bin/bash

# Login credentials
User_Name="TestUser"
PASSWORD="Pa\$\$w0rd"

# Send login request
LOGIN_RESPONSE=$(curl -s -X POST http://localhost:8832/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d "{\"username\":\"$User_Name\", \"password\":\"$PASSWORD\"}")

# Extract certificate from login response
CERTIFICATE=$(echo "$LOGIN_RESPONSE" | jq '.certificate')

# Check if certificate is extracted
if [ -z "$CERTIFICATE" ]; then
  echo "Failed to extract certificate from login response"
  exit 1
fi
echo "Certificate: $CERTIFICATE"
# Send certificate to verify-certificate endpoint
VERIFY_RESPONSE=$(curl -s -X GET http://localhost:8832/api/v1/auth/verify-certificate \
  -H "Content-Type: application/json" \
  -d "$CERTIFICATE")

# Print verification response
echo "Verification Response: $VERIFY_RESPONSE"