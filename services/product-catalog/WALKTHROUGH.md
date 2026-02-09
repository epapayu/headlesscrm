# Verification: TMF620 Product Catalog Service

## Overview
This walkthrough demonstrates how to run and verify the **Product Catalog Service** (TMF620). The service provides REST APIs to browse product offerings.

## 1. Running the Service
Navigate to the service directory and run the application:
```bash
go run .
```
*The service will start on port `8083`.*

## 2. Verification Steps

### Step 2.1: List Product Offerings (GET)
Retrieve all available products.

**Request:**
```bash
curl -s http://localhost:8083/tmf-api/productCatalogManagement/v4/productOffering | jq .
```

**Expected Response:**
- JSON array containing "50GB Freedom Data" and "Unlimited WhatsApp".

### Step 2.2: Get Product Offering (GET)
Retrieve details for "50GB Freedom Data".

**Request:**
```bash
curl -s http://localhost:8083/tmf-api/productCatalogManagement/v4/productOffering/offer-001 | jq .
```

**Expected Response:**
- JSON object with price `100000 IDR`.
