# Verification: TMF622 Product Order Service

## Overview
This walkthrough demonstrates how to run and verify the **Product Order Service** (TMF622). The service handles order capture and lifecycle management.

## 1. Running the Service
Navigate to the service directory and run the application:
```bash
go run .
```
*The service will start on port `8084`.*

## 2. Verification Steps

### Step 2.1: Create Product Order (POST)
Place an order for "50GB Freedom Data".

**Request:**
```bash
curl -X POST http://localhost:8084/tmf-api/productOrderManagement/v4/productOrder \
-H "Content-Type: application/json" \
-d '{
    "externalId": "ext-order-001",
    "description": "New Order for 50GB Data",
    "productOrderItem": [
        {
            "action": "add",
            "quantity": 1,
            "productOffering": {"id": "offer-001", "name": "50GB Freedom Data"}
        }
    ],
    "relatedParty": [{"id": "cust-001", "role": "Customer"}]
}' | jq .
```

**Expected Response:**
- `status`: "Acknowledged"
- `id`: Auto-generated (e.g., `ord-17706...`)

### Step 2.2: Verify Lifecycle (GET)
Wait 3 seconds and check status again (Mock Orchestration).

**Request:**
```bash
curl -s http://localhost:8084/tmf-api/productOrderManagement/v4/productOrder/<ORDER_ID> | jq .
```

**Expected Response:**
- `status`: "InProgress"
