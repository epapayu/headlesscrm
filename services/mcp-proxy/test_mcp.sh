#!/bin/bash
# Send a JSON-RPC tools/list request to the MCP Proxy
./mcp-proxy <<EOF
{"jsonrpc":"2.0","method":"tools/list","id":1}
EOF
