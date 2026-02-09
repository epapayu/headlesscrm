# Verification: KYC Service

## Overview
This walkthrough demonstrates how to run and verify the **KYC Service** (Indonesian Regulatory). The service validates NIK (National ID) and KK (Family Card).

## 1. Running the Service
Navigate to the service directory and run the application:
```bash
go run .
```
*The service will start on port `8082`.*

## 2. Verification Steps

### Step 2.1: Validate Valid Identity (POST)
NIK starts with `3201` (West Java).

**Request:**
```bash
curl -X POST http://localhost:8082/api/kyc/v1/validate \
-H "Content-Type: application/json" \
-d '{"nik": "3201123456789001", "kk": "3201987654321001"}' | jq .
```

**Expected Response:**
```json
{
  "isValid": true,
  "message": "Identity Verified",
  "fullName": "Mock Citizen Name"
}
```

### Step 2.2: Validate Invalid Identity (POST)
Any other NIK.

**Request:**
```bash
curl -X POST http://localhost:8082/api/kyc/v1/validate \
-H "Content-Type: application/json" \
-d '{"nik": "1234567890123456", "kk": "3201987654321001"}' | jq .
```

**Expected Response:**
```json
{
  "isValid": false,
  "message": "Identity Not Found in Dukcapil"
}
```
