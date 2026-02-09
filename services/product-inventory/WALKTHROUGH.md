# Verification: TMF637 Product Inventory Service

## Overview
This walkthrough demonstrates how to run and verify the **Product Inventory Service** (TMF637). The service manages the customer's installed base of products.

## 1. Running the Service
Navigate to the service directory and run the application:
```bash
go run .
```
*The service will start on port `8086`.*

## 2. Verification Steps

### Step 2.1: List Products for Customer (GET)
Retrieve all active products for `cust-001`.

**Request:**
```bash
curl -s "http://localhost:8086/tmf-api/productInventoryManagement/v4/product?relatedParty.id=cust-001" | jq .
```

**Expected Response:**
- JSON array containing "50GB Freedom Data".
- `status`: "Active"
