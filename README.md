# GoContextManager
## About
The `gocontextmanager` is a lightweight go package to help me manage contexts for my chatbots.

## Installation
to install the `gocontextmanager`

```bash
go get github.com/StevenDStanton/gocontextmanager
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/yourusername/gocontextmanager"
)

func main() {
	// Create a new ContextManager instance
	cm := gocontextmanager.NewContextManager()

	// Add messages to a context
	cm.AddContext("testID", "Message 1")
	cm.AddContext("testID", "Message 2")

	// Get messages for a context
	messages := cm.GetContext("testID")
	fmt.Println(messages)
}

```

## Testing

```bash
go test
```

## License 
This package is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.