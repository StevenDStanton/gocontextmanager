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
	"github.com/StevenDStanton/gocontextmanager" // ensure the import is correct
)

func main() {
	cm := gocontextmanager.NewContextManager()

	// Adding messages to a specific context with additional user details
	cm.AddContext("testID", "Hello, World!", "user", "u123", "GlobalUser123", "John Doe")
	cm.AddContext("testID", "How are you today?", "assistant", "a456", "GlobalAssistant", "ChatBot")

	// Retrieving messages from a context
	messages := cm.GetContext("testID")
	for _, msg := range messages {
		fmt.Printf("Message: %s, Role: %s, UserID: %s, UserName: %s\n", msg.content, msg.role, msg.userID, msg.userName)
	}
}

```

## Update

Update the require

```go
require (
    github.com/yourusername/gocontextmanager/v2 v2.0.0
)
```

Then run

```bash
go mod tidy
go mod verify
```

## Testing

```bash
go test
```

## License 
This package is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.