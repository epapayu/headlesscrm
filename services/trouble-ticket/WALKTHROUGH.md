# Walkthrough: Trouble Ticket Service (TMF621)

## Overview
The **Trouble Ticket Service** manages the lifecycle of customer complaints and technical issues. It implements the **TMF621** Open API standard.

## 1. Running the Service
```bash
cd services/trouble-ticket
go run .
```
The service listens on **port 8089**.

## 2. API Endpoints

### List Tickets
```bash
curl http://localhost:8089/tmf-api/troubleTicketManagement/v4/troubleTicket
```

### Create Ticket
```bash
curl -X POST http://localhost:8089/tmf-api/troubleTicketManagement/v4/troubleTicket \
  -H "Content-Type: application/json" \
  -d '{
    "description": "Slow Internet",
    "severity": "Major",
    "type": "NetworkFault",
    "relatedParty": [{"id": "C-001", "role": "Customer", "name": "Budi"}]
  }'
```

### Update Ticket (Patch)
```bash
curl -X PATCH http://localhost:8089/tmf-api/troubleTicketManagement/v4/troubleTicket/{TICKET_ID} \
  -H "Content-Type: application/json" \
  -d '{
    "status": "Resolved",
    "statusChangeReason": "Fixed by technicians",
    "note": [{"author": "Tech", "text": "Replaced faulty router"}]
  }'
```

## 3. Key Features
- **Ticket Lifecycle**: Submitted -> InProgress -> Resolved -> Closed.
- **Related Party**: Links tickets to Customers or Partners.
- **Notes**: Append-only log of actions/comments.
