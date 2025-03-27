package greetings

import (
	"context"

	"github.com/go-tapd/mcp/internal/tools"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct{}

var _ tools.Tool = (*Tool)(nil)

func (t *Tool) Tool() mcp.Tool {
	return mcp.NewTool("tapd greetings",
		mcp.WithDescription("Tapd greetings"),
	)
}

func (t *Tool) Run(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText("Hello, Tapd!"), nil
}
