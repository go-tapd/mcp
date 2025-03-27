package main

import (
	"log"
	"os"

	"github.com/go-tapd/mcp"
	"github.com/go-tapd/tapd"
)

func init() {
	requiredEnvs("TAPD_USERNAME", "TAPD_PASSWORD")
}

func main() {
	var (
		username = os.Getenv("TAPD_USERNAME")
		password = os.Getenv("TAPD_PASSWORD")
	)

	if username == "" || password == "" {
		log.Fatal("missing TAPD_USERNAME or TAPD_PASSWORD")
	}

	client, err := tapd.NewClient(username, password)
	if err != nil {
		log.Fatal(err)
	}

	srv, err := mcp.NewServer(client)
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
