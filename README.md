# Tapd MCP Server

## 📥 Usage

### Use STDIO Server

**Install the package**

```bash
go install github.com/go-tapd/mcp/cmd/tapd-mcp-server@latest
```

**Configure the MCP server**

Below is a configuration example based on Cline, with different configurations for various MCP Clients.

```json
{
  "mcpServers": {
    "github.com/go-tapd/mcp": {
      "command": "{path}/tapd-mcp-server",
      "env": {
        "TAPD_USERNAME": "<YOUR_USERNAME>",
        "TAPD_PASSWORD": "<YOUR_PASSWORD>",
        "TAPD_WORKSPACE_ID": "<YOUR_WORKSPACE_ID>"
      }
    }
  }
}
```

### Use SSE Server

**Install the package**

```bash
go get github.com/go-tapd/mcp
```

**Create a server**

```go
package main

import (
	"log"
	"net/http"

	"github.com/go-tapd/mcp"
	"github.com/go-tapd/tapd"
)

func main() {
	client, err := tapd.NewClient("username", "password")
	if err != nil {
		log.Fatal(err)
	}

	workspaceID := 123456 // replace with your workspace ID

	srv, err := mcp.NewServer(workspaceID, client)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", srv.ServeHTTP)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
```

Visit http://localhost:8080/sse to get the SSE stream.

## 📦 Features

- [x] [获取项目角色ID对照关系](https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/user/get_roles.html)

## 📄 License

[MIT](LICENSE)