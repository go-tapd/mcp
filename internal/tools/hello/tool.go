package hello

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-tapd/mcp/internal/tools"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct{}

var _ tools.Tool = (*Tool)(nil)

func (t *Tool) Tool() mcp.Tool {
	return mcp.NewTool("hello_world",
		mcp.WithDescription("Say hello to someone"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the person to greet"),
		),
	)
}

func (t *Tool) Run(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, ok := request.Params.Arguments["name"].(string)
	if !ok {
		return nil, errors.New("name must be a string")
	}

	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}
