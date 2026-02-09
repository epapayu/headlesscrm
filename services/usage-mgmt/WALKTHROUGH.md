# Verification: TMF635 Usage Management Service

## Overview
This walkthrough demonstrates how to run and verify the **Usage Management Service** (TMF635). The service handles usage records (IPDR/CDR) and simple rating.

## 1. Running the Service
Navigate to the service directory and run the application:
```bash
go run .
```
*The service will start on port `8085`.*

## 2. Verification Steps

### Step 2.1: Create & Rate Usage (POST)
Simulate 500MB Data usage.

**Request:**
```bash
curl -X POST http://localhost:8085/tmf-api/usageManagement/v4/usage \
-H "Content-Type: application/json" \
-d '{
    "type": "Data",
    "description": "Video Streaming",
    "usageCharacteristic": [{"name": "volume_mb", "value": "500"}],
    "relatedParty": [{"id": "cust-001", "role": "Customer"}]
}' | jq .
```

**Expected Response:**
- `status`: "Rated"
- `ratedProductUsage`: Contains `5000 IDR` (assuming 10 IDR/MB).
