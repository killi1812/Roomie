#!/bin/bash

# Login credentials
User_Name="TestUser"
PASSWORD="Pa\$\$w0rd"

# Send login request
LOGIN_RESPONSE=$(curl -sk -X POST https://localhost:8832/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d "{\"username\":\"$User_Name\", \"password\":\"$PASSWORD\"}")

# Extract certificate from login response
CERTIFICATE=$(echo "$LOGIN_RESPONSE" | jq '.certificate')

echo "response $LOGIN_RESPONSE"
# Check if certificate is extracted
if [ -z "$CERTIFICATE" ]; then
  echo "Failed to extract certificate from login response"
  exit 1
fi
echo "Certificate: $CERTIFICATE"
# Send certificate to verify-certificate endpoint
VERIFY_RESPONSE=$(curl -ks -X GET https://localhost:8832/api/v1/auth/verify-certificate \
  -H "Content-Type: application/json" \
  -d "$CERTIFICATE")

# Print verification response
echo "Verification Response: $VERIFY_RESPONSE"