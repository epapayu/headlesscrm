# Verification: TMF639 Resource Inventory Service

## Overview
This walkthrough demonstrates how to run and verify the **Resource Inventory Service** (TMF639). The service manages physical (SIM) and logical (MSISDN) resources.

## 1. Running the Service
Navigate to the service directory and run the application:
```bash
go run .
```
*The service will start on port `8088`.*

## 2. Verification Steps

### Step 2.1: List MSISDN Resources (GET)
Retrieve all available MSISDNs.

**Request:**
```bash
curl -s "http://localhost:8088/tmf-api/resourceInventoryManagement/v4/resource?category=MSISDN" | jq .
```

**Expected Response:**
- JSON array containing "628123456789" and "628129999999".
- `category`: "MSISDN"
