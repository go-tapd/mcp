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
