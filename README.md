# Volumez REST Client

This is a Go client library for the Volumez API, generated from the OpenAPI specification.

## Installation

```bash
go get github.com/VolumezTech/volumez-rest-client
```

## Authentication

The client requires an ID token for authentication. Pass it as an Authorization header to each API call:

```go
client := openapi.NewAPIClient(config)

// Create a context
ctx := context.Background()

// Example: List volumes with authentication
volumes, _, err := client.VolumesAPI.VolumesList(ctx).Authorization("YOUR_ID_TOKEN").Execute()
if err != nil {
    log.Fatalf("Error listing volumes: %v", err)
}


```

## Examples

See the `examples` directory for more detailed examples of using the client.

## License

This client is generated from the Volumez OpenAPI specification and is subject to the same license terms as the Volumez API.

