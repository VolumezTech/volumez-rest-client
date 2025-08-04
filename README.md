# Volumez OpenAPI Client

This is a Go client library for the Volumez API, generated from the OpenAPI specification.

## Installation

```bash
go get 	"github.com/VolumezTech/volumez-rest-client/pkg/openapi"

```

## Usage

Here's a basic example of how to use the client:

```go
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

    // Set your API key or token
    config.AddDefaultHeader("Authorization", "Bearer YOUR_API_TOKEN")

    // Create the client
    client := openapi.NewAPIClient(config)

    // Create a context
    ctx := context.Background()

    // List volumes
    volumes, _, err := client.VolumesAPI.VolumesList(ctx).Execute()
    if err != nil {
        log.Fatalf("Error listing volumes: %v", err)
    }
    fmt.Printf("Found %d volumes\n", len(volumes))
}
```

## Available APIs

The client provides access to the following APIs:

- VolumesAPI: Manage volumes
- NodesAPI: Manage nodes
- SystemAPI: System information and features
- And more...

## Authentication

The client requires an API token for authentication. Set it in the configuration:

```go
config := openapi.NewConfiguration()
config.AddDefaultHeader("Authorization", "Bearer YOUR_API_TOKEN")
```

## Error Handling

All API methods return three values:
1. The response data
2. The HTTP response
3. An error (if any)

Always check the error before using the response data:

```go
response, _, err := client.SomeAPI.SomeMethod(ctx).Execute()
if err != nil {
    log.Fatalf("Error: %v", err)
}
```

## Examples

See the `examples` directory for more detailed examples of using the client.

## License

This client is generated from the Volumez OpenAPI specification and is subject to the same license terms as the Volumez API.

## Development

### Prerequisites

- Go 1.18 or later
- Make
- [Squire](https://github.com/your-org/squire) (for OpenAPI client generation)

### Building

The client library is generated from the OpenAPI specification. To regenerate the client:

```bash
make generate-openapi-client
```

### Validation

To validate the generated client:

```bash
make validate-openapi-client
```

### Cleaning

To clean the generated files:

```bash
make clean
```

## Available Make Commands

- `make generate-openapi-client`: Generates the OpenAPI client from the specification
- `make validate-openapi-client`: Validates the generated client
- `make clean`: Removes generated files
