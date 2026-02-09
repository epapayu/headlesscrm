package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Service URLs
const (
	BalanceServiceURL   = "http://localhost:8080/tmf-api/accountBalanceManagement/v4"
	CustomerServiceURL  = "http://localhost:8081/tmf-api/customerManagement/v4"
	KYCServiceURL       = "http://localhost:8082/api/kyc/v1"
	CatalogServiceURL   = "http://localhost:8083/tmf-api/productCatalogManagement/v4"
	OrderServiceURL     = "http://localhost:8084/tmf-api/productOrderManagement/v4"
	UsageServiceURL     = "http://localhost:8085/tmf-api/usageManagement/v4"
	InventoryServiceURL = "http://localhost:8086/tmf-api/productInventoryManagement/v4"
	ServiceOrderURL     = "http://localhost:8087/tmf-api/serviceOrderManagement/v4"
	ResourceServiceURL  = "http://localhost:8088/tmf-api/resourceInventoryManagement/v4"
)

func main() {
	s := server.NewMCPServer(
		"Headless CRM MCP",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	// --- Account Balance (TMF654) ---
	s.AddTool(mcp.NewTool("get_balance",
		mcp.WithDescription("Get account balance"),
		mcp.WithString("id", mcp.Required(), mcp.Description("Wallet ID")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("invalid arguments"), nil
		}
		id, _ := args["id"].(string)
		return httpGet(fmt.Sprintf("%s/balance/%s", BalanceServiceURL, id))
	})

	s.AddTool(mcp.NewTool("adjust_balance",
		mcp.WithDescription("Credit or Debit a wallet"),
		mcp.WithString("wallet_id", mcp.Required(), mcp.Description("Wallet ID")),
		mcp.WithNumber("amount", mcp.Required(), mcp.Description("Amount")),
		mcp.WithString("unit", mcp.Required(), mcp.Description("Currency Unit (e.g. IDR)")),
		mcp.WithString("type", mcp.Required(), mcp.Description("adjustment or deduction")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("invalid arguments"), nil
		}
		payload := map[string]interface{}{
			"relatedWallet": map[string]interface{}{"id": args["wallet_id"]},
			"amount":        map[string]interface{}{"value": args["amount"], "unit": args["unit"]},
			"type":          args["type"],
			"description":   "MCP Adjustment",
		}
		return httpPost(fmt.Sprintf("%s/balanceAdjustment", BalanceServiceURL), payload)
	})

	// --- Customer Mgmt (TMF629) ---
	s.AddTool(mcp.NewTool("get_customer",
		mcp.WithDescription("Get customer profile"),
		mcp.WithString("id", mcp.Required(), mcp.Description("Customer ID")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("invalid arguments"), nil
		}
		id, _ := args["id"].(string)
		return httpGet(fmt.Sprintf("%s/customer/%s", CustomerServiceURL, id))
	})

	s.AddTool(mcp.NewTool("create_customer",
		mcp.WithDescription("Create new customer"),
		mcp.WithString("name", mcp.Required(), mcp.Description("Full Name")),
		mcp.WithString("email", mcp.Required(), mcp.Description("Email Address")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("invalid arguments"), nil
		}
		payload := map[string]interface{}{
			"name": args["name"],
			"contactMedium": []map[string]interface{}{
				{
					"mediumType": "Email",
					"preferred":  true,
					"characteristic": map[string]interface{}{
						"emailAddress": args["email"],
					},
				},
			},
		}
		return httpPost(fmt.Sprintf("%s/customer", CustomerServiceURL), payload)
	})

	// --- KYC Service ---
	s.AddTool(mcp.NewTool("validate_identity",
		mcp.WithDescription("Validate NIK and KK (Dukcapil)"),
		mcp.WithString("nik", mcp.Required(), mcp.Description("National ID (16 digits)")),
		mcp.WithString("kk", mcp.Required(), mcp.Description("Family Card ID (16 digits)")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("invalid arguments"), nil
		}
		payload := map[string]interface{}{
			"nik": args["nik"],
			"kk":  args["kk"],
		}
		return httpPost(fmt.Sprintf("%s/validate", KYCServiceURL), payload)
	})

	// --- Product Catalog (TMF620) ---
	s.AddTool(mcp.NewTool("list_products",
		mcp.WithDescription("List available product offerings"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return httpGet(fmt.Sprintf("%s/productOffering", CatalogServiceURL))
	})

	// --- Product Inventory (TMF637) ---
	s.AddTool(mcp.NewTool("get_installed_base",
		mcp.WithDescription("Get products owned by a customer"),
		mcp.WithString("customer_id", mcp.Required(), mcp.Description("Customer ID")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("invalid arguments"), nil
		}
		cid, _ := args["customer_id"].(string)
		return httpGet(fmt.Sprintf("%s/product?relatedParty.id=%s", InventoryServiceURL, cid))
	})

	// --- Product Order (TMF622) ---
	s.AddTool(mcp.NewTool("create_order",
		mcp.WithDescription("Create a new product order"),
		mcp.WithString("customer_id", mcp.Required(), mcp.Description("Customer ID")),
		mcp.WithString("offering_id", mcp.Required(), mcp.Description("Product Offering ID")),
		mcp.WithString("offering_name", mcp.Required(), mcp.Description("Product Offering Name")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("invalid arguments"), nil
		}
		payload := map[string]interface{}{
			"externalId":  fmt.Sprintf("ord-%s", args["customer_id"]),
			"description": "MCP Order",
			"relatedParty": []map[string]interface{}{
				{"id": args["customer_id"], "role": "Customer"},
			},
			"productOrderItem": []map[string]interface{}{
				{
					"action":   "add",
					"quantity": 1,
					"productOffering": map[string]interface{}{
						"id":   args["offering_id"],
						"name": args["offering_name"],
					},
				},
			},
		}
		return httpPost(fmt.Sprintf("%s/productOrder", OrderServiceURL), payload)
	})

	s.AddTool(mcp.NewTool("get_order",
		mcp.WithDescription("Get order status"),
		mcp.WithString("id", mcp.Required(), mcp.Description("Order ID")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("invalid arguments"), nil
		}
		id, _ := args["id"].(string)
		return httpGet(fmt.Sprintf("%s/productOrder/%s", OrderServiceURL, id))
	})

	// --- Usage Mgmt (TMF635) ---
	s.AddTool(mcp.NewTool("report_usage",
		mcp.WithDescription("Report usage (Data/Voice/SMS)"),
		mcp.WithString("type", mcp.Required(), mcp.Description("Usage Type (Data/Voice/SMS)")),
		mcp.WithString("customer_id", mcp.Required(), mcp.Description("Customer ID")),
		mcp.WithString("quantity", mcp.Required(), mcp.Description("Quantity (e.g. 500 for 500MB)")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("invalid arguments"), nil
		}
		charName := "volume_mb"
		if args["type"] == "Voice" {
			charName = "duration_sec"
		} else if args["type"] == "SMS" {
			charName = "count"
		}
		payload := map[string]interface{}{
			"type":        args["type"],
			"description": "MCP Usage",
			"relatedParty": []map[string]interface{}{
				{"id": args["customer_id"], "role": "Customer"},
			},
			"usageCharacteristic": []map[string]interface{}{
				{"name": charName, "value": args["quantity"]},
			},
		}
		return httpPost(fmt.Sprintf("%s/usage", UsageServiceURL), payload)
	})

    // --- Resource Inventory (TMF639) ---
	s.AddTool(mcp.NewTool("list_resources",
		mcp.WithDescription("List resources by category"),
		mcp.WithString("category", mcp.Required(), mcp.Description("Category (MSISDN/SIM)")),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("invalid arguments"), nil
		}
		cat, _ := args["category"].(string)
		return httpGet(fmt.Sprintf("%s/resource?category=%s", ResourceServiceURL, cat))
	})

	// Start server
	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
	}
}

// Helpers

func httpGet(url string) (*mcp.CallToolResult, error) {
	resp, err := http.Get(url)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to call service: %v", err)), nil
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return mcp.NewToolResultText(string(body)), nil
}

func httpPost(url string, payload interface{}) (*mcp.CallToolResult, error) {
	jsonPayload, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to call service: %v", err)), nil
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return mcp.NewToolResultText(string(body)), nil
}
