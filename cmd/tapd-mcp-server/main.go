package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-tapd/mcp"
	"github.com/go-tapd/tapd"
)

func init() {
	requiredEnvs("TAPD_USERNAME", "TAPD_PASSWORD", "TAPD_WORKSPACE_ID")
}

func main() {
	var (
		username  = os.Getenv("TAPD_USERNAME")
		password  = os.Getenv("TAPD_PASSWORD")
		workspace = os.Getenv("TAPD_WORKSPACE_ID")
	)

	if username == "" || password == "" || workspace == "" {
		log.Fatal("missing TAPD_USERNAME, TAPD_PASSWORD or TAPD_WORKSPACE_ID")
	}

	workspaceID, err := convertToInt(workspace)
	if err != nil {
		log.Fatalf("invalid TAPD_WORKSPACE_ID: %s", err)
	}

	client, err := tapd.NewClient(username, password)
	if err != nil {
		log.Fatal(err)
	}

	srv, err := mcp.NewServer(workspaceID, client)
	if err != nil {
		log.Fatal(err)
	}

	if err := srv.ServerStdio(); err != nil {
		log.Fatal(err)
	}
}

func requiredEnvs(keys ...string) {
	var missing []string
	for _, key := range keys {
		if _, ok := os.LookupEnv(key); !ok {
			missing = append(missing, key)
		}
	}
	if len(missing) > 0 {
		log.Fatalf("missing required env vars: %v", missing)
	}
}

func convertToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid int: %s", s)
	}
	return i, nil
}
