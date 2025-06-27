package main

import (
	"context"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"os"
)

func main() {
	fmt.Printf("Starting hello world from Darp app\n")
	err := os.Setenv("DAPR_CLIENT_TIMEOUT_SECONDS", "2")
	if err != nil {
		fmt.Printf("💥💥 error setting DAPR_CLIENT_TIMEOUT_SECONDS environment variable: %s", err)
		os.Exit(1)
	}
	client, err := dapr.NewClient()
	if err != nil {
		fmt.Printf("💥💥 error creating Dapr client: %s\n", err)
		fmt.Println("Did you run the Dapr side car in another console window ?")
		fmt.Println("dapr run --app-id myapp --app-port 8080 --dapr-http-port 3500 --dapr-grpc-port 50001")
		os.Exit(1)
	}
	defer client.Close()

	// Get Dapr metadata to retrieve the version
	metadata, err := client.GetMetadata(context.Background())
	if err != nil {
		fmt.Printf("💥💥 error getting Dapr metadata: %s\n", err)
		os.Exit(1)
	}

	// Print the entire metadata response to inspect available fields
	fmt.Printf("Dapr Metadata: %+v\n", metadata)
	// try to retrieve and print the Dapr runtime version
	if version, ok := metadata.ExtendedMetadata["daprRuntimeVersion"]; ok {
		fmt.Printf("Dapr Runtime Version: %s\n", version)
	} else {
		fmt.Println("💥💥 Dapr runtime version not found in metadata")
	}

}
