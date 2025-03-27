package mcp

import (
	"log"
	"net/http"

	"github.com/go-tapd/mcp/internal/tools"
	"github.com/go-tapd/mcp/internal/tools/hello"
	"github.com/go-tapd/tapd"
	"github.com/mark3labs/mcp-go/server"
)

type Server struct {
	mcpServer  *server.MCPServer
	tapdClient *tapd.Client
}

var _ http.Handler = (*Server)(nil)

func NewServer(client *tapd.Client, opts ...Option) (*Server, error) {
	o, err := newOptions(opts...)
	if err != nil {
		return nil, err
	}

	srv := &Server{
		tapdClient: client,
		mcpServer:  server.NewMCPServer(o.name, Version()),
	}

	srv.registerTools()

	return srv, nil
}

func (s *Server) registerTools() {
	tools.RegisterTools(s.mcpServer,
		&hello.Tool{},
	)
}

func (s *Server) ServerStdio() error {
	log.Println("Tapd MCP server is running")
	return server.ServeStdio(s.mcpServer)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Tapd MCP server is running")
	server.NewSSEServer(s.mcpServer).ServeHTTP(w, r)
}
