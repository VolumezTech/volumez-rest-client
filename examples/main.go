package main

import (
	"context"
	"fmt"
	"log"

	"github.com/VolumezTech/volumez-rest-client/pkg/openapi"
)

func main() {
	// Create a new client configuration
	config := openapi.NewConfiguration()
	config.Host = "https://api.volumez.com"

	// Create the client
	client := openapi.NewAPIClient(config)

	// Create a context
	ctx := context.Background()

	// List volumes
	volumes, _, err := client.VolumesAPI.VolumesList(ctx).Authorization("CLIENT_ID_TOKEN").Execute()
	if err != nil {
		log.Fatalf("Error listing volumes: %v", err)
	}
	fmt.Printf("Found %d volumes\n", len(volumes))
}
