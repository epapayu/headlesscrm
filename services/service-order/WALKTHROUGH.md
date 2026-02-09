# Verification: TMF641 Service Order Service

## Overview
This walkthrough demonstrates how to run and verify the **Service Order Service** (TMF641). The service handles technical provisioning requests (e.g., from Product Order Decomposition).

## 1. Running the Service
Navigate to the service directory and run the application:
```bash
go run .
```
*The service will start on port `8087`.*

## 2. Verification Steps

### Step 2.1: Create Service Order (POST)
Provision a "Mobile Data Service".

**Request:**
```bash
curl -X POST http://localhost:8087/tmf-api/serviceOrderManagement/v4/serviceOrder \
-H "Content-Type: application/json" \
-d '{
    "externalId": "ext-srv-001",
    "description": "Provision 50GB Data",
    "serviceOrderItem": [
        {
            "action": "add",
            "service": {
                "name": "Mobile Data Service",
                "serviceCharacteristic": [{"name": "quota", "value": "50GB"}]
            }
        }
    ]
}' | jq .
```

**Expected Response:**
- `status`: "Acknowledged"
- `id`: Auto-generated (e.g., `srv-ord-17706...`)

### Step 2.2: Verify Provisioning (GET)
Wait 3 seconds and check status again (Mock Provisioning).

**Request:**
```bash
curl -s http://localhost:8087/tmf-api/serviceOrderManagement/v4/serviceOrder/<ORDER_ID> | jq .
```

**Expected Response:**
- `status`: "Completed"
