# Verification: TMF654 Account Balance Service

## Overview
This walkthrough demonstrates how to run and verify the **Account Balance Service** (TMF654). The service provides REST APIs to query balances and perform adjustments (credit/debit).

## 1. Prerequisites
- **Go 1.20+** installed.
- **Curl** for testing endpoints.

## 2. Running the Service
Navigate to the service directory and run the application:
```bash
go run .
```
*The service will start on port `8080`.*

## 3. Verification Steps

### Step 3.1: Check Balance (GET)
Retrieve the balance for the mock wallet `bal-123`.

**Request:**
```bash
curl -s http://localhost:8080/tmf-api/accountBalanceManagement/v4/balance/bal-123 | jq .
```

**Expected Response:**
```json
{
  "id": "bal-123",
  "status": "Active",
  "amount": {
    "value": 50000,
    "unit": "IDR"
  },
  "validFor": { ... }
}
```

### Step 3.2: Adjust Balance (Top-up/Credit) (POST)
Add `5000 IDR` to the wallet.

**Request:**
```bash
curl -X POST http://localhost:8080/tmf-api/accountBalanceManagement/v4/balanceAdjustment \
-H "Content-Type: application/json" \
-d '{
    "type": "adjustment",
    "amount": {"value": 5000, "unit": "IDR"},
    "description": "Bonus Credit",
    "reasonCode": "BONUS_001",
    "relatedWallet": {"id": "bal-123"}
}' | jq .
```

**Expected Response:**
- HTTP Status: `201 Created`
- Body: JSON object of the created adjustment with a new `id`.

### Step 3.3: Verify New Balance (GET)
Repeat Step 3.1. The balance should now be `55000`.

## 4. Troubleshooting
- **Port In Use**: If port 8080 is busy, modify `main.go` to use a different port.
- **Go Mod Errors**: Run `go mod tidy` if you encounter dependency issues.
