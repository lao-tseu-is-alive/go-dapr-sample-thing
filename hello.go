package main

import (
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"os"
)

func main() {
	err := os.Setenv("DAPR_CLIENT_TIMEOUT_SECONDS", "3")
	if err != nil {
		fmt.Printf("error setting DAPR_CLIENT_TIMEOUT_SECONDS environment variable: %s", err)
		return
	}
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()
}
