# Verification: MCP Proxy Service

## Overview
This walkthrough demonstrates the **MCP Proxy Service**. This service acts as a **Model Context Protocol (MCP)** server, exposing all Headless CRM microservices (TMF APIs) as tools to AI agents (e.g., Claude Desktop, cursor, or your own agent).

## 1. Prerequisites
- Go 1.20+ installed.
- All other microservices running (ports 8080-8088).

## 2. Running the MCP Server
You can run the server directly using `go run`:
```bash
cd services/mcp-proxy
go run .
```
*Note: It communicates via Stdio (Standard Input/Output) adhering to JSON-RPC 2.0.*

## 3. Configuring an MCP Client
 To use this with an MCP-compliant client (like Claude Desktop), add the following to your configuration:

```json
{
  "mcpServers": {
    "headless-crm": {
      "command": "go",
      "args": ["run", "/absolute/path/to/headless-crm/services/mcp-proxy"]
    }
  }
}
```

## 4. Available Tools
The proxy exposes the following tools:
- `get_balance(id)`
- `adjust_balance(wallet_id, amount, unit, type)`
- `get_customer(id)`
- `create_customer(name, email)`
- `validate_identity(nik, kk)`
- `list_products()`
- `get_installed_base(customer_id)`
- `create_order(customer_id, offering_id, offering_name)`
- `get_order(id)`
- `report_usage(type, customer_id, quantity)`
- `list_resources(category)`

## 5. Manual Verification (JSON-RPC)
You can test it manually by sending a JSON-RPC request:

```bash
./mcp-proxy <<EOF
{"jsonrpc":"2.0","method":"tools/list","id":1}
EOF
```
