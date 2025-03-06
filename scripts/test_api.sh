#!/bin/bash

BASE_URL="http://localhost:8080"

echo "1. Checking health..."
curl -s "$BASE_URL/health" | jq

echo "2. Fetching accounts..."
curl -s "$BASE_URL/accounts" | jq

echo "3. Performing transfer (Mark -> Jane, 20)..."
curl -s -X POST "$BASE_URL/transfer" \
    -H "Content-Type: application/json" \
    -d '{"from":"Mark","to":"Jane","amount":20}' | jq

echo "4. Fetching accounts after transfer..."
curl -s "$BASE_URL/accounts" | jq
