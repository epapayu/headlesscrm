# Verification: TMF629 Customer Management Service

## Overview
This walkthrough demonstrates how to run and verify the **Customer Management Service** (TMF629). The service provides REST APIs to manage customer profiles.

## 1. Running the Service
Navigate to the service directory and run the application:
```bash
go run .
```
*The service will start on port `8081`.*

## 2. Verification Steps

### Step 2.1: Create Customer (POST)
Create a new customer `Siti Aminah`.

**Request:**
```bash
curl -X POST http://localhost:8081/tmf-api/customerManagement/v4/customer \
-H "Content-Type: application/json" \
-d '{
    "name": "Siti Aminah",
    "contactMedium": [
        {
            "mediumType": "Email",
            "preferred": true,
            "characteristic": {"emailAddress": "siti@example.com"}
        }
    ]
}' | jq .
```

**Expected Response:**
- HTTP Status: `201 Created`
- `status`: "Initial" (Default logic)
- `id`: Auto-generated (e.g., `cust-1739086...`)

### Step 2.2: Get Customer (GET)
Use the `id` from the previous step (or use mock `cust-001`).

**Request:**
```bash
curl -s http://localhost:8081/tmf-api/customerManagement/v4/customer/cust-001 | jq .
```

**Expected Response:**
- JSON object for "Budi Santoso".
