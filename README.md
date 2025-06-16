# Environment Module

[![Tests Status](https://github.com/dracory/env/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/dracory/env/actions/workflows/tests.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/dracory/env.svg)](https://pkg.go.dev/github.com/dracory/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/dracory/env)](https://goreportcard.com/report/github.com/dracory/env)

A Go module for managing environment variables with support for secure environment variable handling and .env file loading.

## Features

- Load environment variables from `.env` files
- Secure environment variable encryption/decryption
- Simple and intuitive API
- Type-safe environment variable access
- Support for default values


## API Reference

### Functions

- `Initialize(envFilePath ...string)` - Initialize environment variables from .env files. Can accept multiple file paths.
- `Value(key string) string` - Get an environment variable. Returns empty string if not set.
- `Must(key string) string` - Get a required environment variable. Panics if the variable is not set.
- `Get(key, defaultValue string) string` - Get environment variable with default value.
- `Set(key, value string) error` - Set an environment variable.

### Advanced Functions

- `VaultLoad(options struct{...}) error` - Load environment variables from an encrypted vault file or a string.
- `Encrypt(data, key string) (string, error)` - Encrypt sensitive data.
- `Decrypt(encryptedData, key string) (string, error)` - Decrypt data.

## Installation

```bash
go get github.com/dracory/env
```

## Usage

### Working with .env Files

Create a `.env` file in your project root:

```env
DB_HOST="localhost"
DB_PORT="5432"
SECRET_KEY="your-secret-key"
```

### Basic Usage

```go
package main

import (
	"fmt"
	"github.com/dracory/env"
)

func main() {
	// Initialize environment variables from .env files
	env.Initialize(".env")
	
	// Get environment variable with default value
	dbHost := env.Get("DB_HOST", "localhost")
	
	// Get required environment variable (panics if not set)
	dbPort := env.MustGet("DB_PORT")
	
	fmt.Printf("Database: %s:%s\n", dbHost, dbPort)
}
```


### Advanced. Env Vault Loading
Env vault loading is a feature that allows you to load environment variables
from an encrypted vault file or a string.

```go
// Load environment variables from an encrypted vault file
err := env.VaultLoad(struct {
    Password      string
    VaultFilePath string
    VaultContent  string
} {
    Password:     "your-password",
    VaultFilePath: ".env.vault",
})
if err != nil {
    // handle error
}
```

### Advanced. Environment Variable Encryption

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

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is dual-licensed under the following terms:

- For non-commercial use, you may choose either the GNU Affero General Public License v3.0 (AGPLv3) or a separate commercial license (see below). You can find a copy of the AGPLv3 at: https://www.gnu.org/licenses/agpl-3.0.txt

- For commercial use, a separate commercial license is required. Commercial licenses are available for various use cases. Please contact me via my contact page at https://lesichkov.co.uk/contact to obtain a commercial license.
