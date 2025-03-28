package roles

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-tapd/mcp/internal/tools"
	"github.com/go-tapd/tapd"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	client *tapd.Client
	tool   mcp.Tool
}

var _ tools.Tool = (*Tool)(nil)

func NewTool(client *tapd.Client) *Tool {
	return &Tool{
		client: client,
		tool: mcp.NewTool("get_user_roles",
			mcp.WithDescription("获取项目角色ID对照关系"),
			mcp.WithString("workspace_id", // todo: 调整为int类型？
				mcp.Required(),
				mcp.Description("项目ID"),
			),
		),
	}
}

func (t *Tool) Tool() mcp.Tool {
	return t.tool
}

func (t *Tool) Run(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	workspaceID, ok := request.Params.Arguments["workspace_id"]
	if !ok {
		return nil, fmt.Errorf("workspace_id is required")
	}
	workspaceIDString, ok := workspaceID.(string)
	if !ok {
		return nil, fmt.Errorf("workspace_id must be a string")
	}
	workspaceIDInt, err := strconv.Atoi(workspaceIDString)
	if err != nil {
		return nil, err
	}

	roles, _, err := t.client.UserService.GetRoles(ctx, &tapd.GetRolesRequest{
		WorkspaceID: tapd.Ptr(workspaceIDInt),
	})
	if err != nil {
		return nil, err
	}

	rolesBytes, err := json.Marshal(roles)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(rolesBytes)), nil
}
