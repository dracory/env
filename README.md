# Environment Module

![tests](https://github.com/dracory/env/workflows/tests/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/dracory/env.svg)](https://pkg.go.dev/github.com/dracory/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/dracory/env)](https://goreportcard.com/report/github.com/dracory/env)

A Go module for managing environment variables with support for secure environment variable handling and .env file loading.

## Features

- Load environment variables from `.env` files
- Secure environment variable encryption/decryption
- Simple and intuitive API
- Type-safe environment variable access
- Support for default values

## Installation

```bash
go get github.com/dracory/env
```

## Usage

### Basic Usage

```go
package main

import (
	"fmt"
	"github.com/dracory/env"
	_ "github.com/joho/godotenv/autoload" // Auto-load .env file
)

func main() {
	// Get environment variable with default value
	dbHost := env.Get("DB_HOST", "localhost")
	
	// Get required environment variable (panics if not set)
	dbPort := env.MustGet("DB_PORT")
	
	fmt.Printf("Database: %s:%s\n", dbHost, dbPort)
}
```

### Working with .env Files

Create a `.env` file in your project root:

```env
DB_HOST="localhost"
DB_PORT="5432"
SECRET_KEY="your-secret-key"
```

### Secure Environment Variables

```go
// Encrypt sensitive data
encrypted, err := env.Encrypt("my-sensitive-data", "encryption-key")
if err != nil {
    // handle error
}

// Decrypt the data
decrypted, err := env.Decrypt(encrypted, "encryption-key")
if err != nil {
    // handle error
}
```

## API Reference

### Functions

- `Get(key, defaultValue string) string` - Get environment variable with default value
- `MustGet(key string) string` - Get required environment variable (panics if not set)
- `Set(key, value string) error` - Set environment variable
- `Encrypt(data, key string) (string, error)` - Encrypt sensitive data
- `Decrypt(encryptedData, key string) (string, error)` - Decrypt data

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
