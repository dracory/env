# Environment Module

[![Tests Status](https://github.com/dracory/env/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/dracory/env/actions/workflows/tests.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/dracory/env.svg)](https://pkg.go.dev/badge/github.com/dracory/env)
[![Go Report Card](https://goreportcard.com/badge/github.com/dracory/env)](https://goreportcard.com/report/github.com/dracory/env)

A Go module for managing environment variables with support for .env loading, value processing (base64/obfuscated), and optional vault-based loading.

## Features

- Load environment variables from `.env` files (`Load`)
- Process values with `base64:` and `obfuscated:` prefixes automatically
- Simple and intuitive API for `string`, `bool`, `int`, and `float64` types.
- Each data type (`String`, `Bool`, `Int`, `Float`) provides four functions for flexible error handling:
    - `Get...`: Returns the value or a zero-value (`"", false, 0`) if not found.
    - `Get...OrDefault`: Returns a specified default value if not found.
    - `Get...OrError`: Returns an error if not found or invalid.
    - `Get...OrPanic`: Panics if not found or invalid.
- Optional encrypted vault loading (`LoadVault`)
- Note: `Float64`-named functions remain available as aliases of `Float` for compatibility.

## API Reference

### Loading Functions

- `Load(envFilePath ...string)` – Load environment variables from `.env` files. Defaults to `.env` and also attempts any additional paths provided.
- `LoadVault(options struct{ Password string; VaultFilePath string; VaultContent string }) error` – Load environment variables from an encrypted vault file or a vault string.

### String Functions

- `GetString(key string) string`
- `GetStringOrDefault(key string, defaultValue string) string`
- `GetStringOrError(key string) (string, error)`
- `GetStringOrPanic(key string) string`

### Bool Functions

- `GetBool(key string) bool`
- `GetBoolOrDefault(key string, defaultValue bool) bool`
- `GetBoolOrError(key string) (bool, error)`
- `GetBoolOrPanic(key string) bool`

### Int Functions

- `GetInt(key string) int`
- `GetIntOrDefault(key string, defaultValue int) int`
- `GetIntOrError(key string) (int, error)`
- `GetIntOrPanic(key string) int`

### Float Functions

- `GetFloat(key string) float64`
- `GetFloatOrDefault(key string, defaultValue float64) float64`
- `GetFloatOrError(key string) (float64, error)`
- `GetFloatOrPanic(key string) float64`

Compatibility: `GetFloat64`, `GetFloat64OrDefault`, `GetFloat64OrError`, and `GetFloat64OrPanic` are available as aliases.

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
API_TIMEOUT="30"
DEBUG_MODE="true"
```

### Basic Usage

```go
package main

import (
	"fmt"
	"github.com/dracory/env"
	"log"
)

func main() {
	// Load environment variables from .env files
	env.Load(".env")

	// Get a string variable with a default value
	dbHost := env.GetStringOrDefault("DB_HOST", "localhost")

	// Get a required string variable (panics if not set)
	secretKey := env.GetStringOrPanic("SECRET_KEY")

	// Get an integer variable, handle error if not found or invalid
	dbPort, err := env.GetIntOrError("DB_PORT")
	if err != nil {
		log.Fatalf("Invalid or missing DB_PORT: %v", err)
	}

	// Get a boolean variable
	debugMode := env.GetBool("DEBUG_MODE")

	fmt.Printf("Database: %s:%d\n", dbHost, dbPort)
	fmt.Printf("Secret Key Loaded: %v\n", secretKey != "")
	fmt.Printf("Debug Mode: %v\n", debugMode)
}
```

### Value Processing
`GetString`, `GetStringOrDefault`, `GetStringOrError`, and `GetStringOrPanic` automatically process these prefixes:

- `base64:<encoded>` – Decodes using URL-safe base64.
- `obfuscated:<text>` – Deobfuscates using `github.com/gouniverse/envenc`.

This lets you safely store encoded/obfuscated values in `.env` or other sources while retrieving plain values at runtime.

### Boolean Parsing
Boolean functions (`GetBool`, etc.) parse values with flexibility:

- **True values**: `"true"`, `"True"`, `"TRUE"`, `"T"`, `"t"`, `"1"`, `"yes"`, `"Yes"`, `"YES"`, and any positive number.
- **False values**: `"false"`, `"False"`, `"FALSE"`, `"F"`, `"f"`, `"0"`, `"no"`, `"No"`, `"NO"`, and any negative number.
- Any other or empty value returns `false` (for `GetBool`) or the specified default.

### Advanced: Env Vault Loading
Env vault loading allows you to load environment variables from an encrypted vault file or a string.

```go
// Load environment variables from an encrypted vault file
er := env.LoadVault(struct {
    Password      string
    VaultFilePath string
    VaultContent  string
}{
    Password:      "your-password",
    VaultFilePath: ".env.vault",
})
if err != nil {
    // handle error
}
```

### Notes on Load
`Load()` will attempt to load from a default `.env` file, and then from any additional file paths you pass in. Missing files are silently skipped.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is dual-licensed under the following terms:

- For non-commercial use, you may choose either the GNU Affero General Public License v3.0 (AGPLv3) or a separate commercial license (see below). You can find a copy of the AGPLv3 at: https://www.gnu.org/licenses/agpl-3.0.txt

- For commercial use, a separate commercial license is required. Commercial licenses are available for various use cases. Please contact me via my contact page at https://lesichkov.co.uk/contact to obtain a commercial license.


## API Reference

### Loading Functions

- `Load(envFilePath ...string)` – Load environment variables from `.env` files. Defaults to `.env` and also attempts any additional paths provided.
- `LoadVault(options struct{ Password string; VaultFilePath string; VaultContent string }) error` – Load environment variables from an encrypted vault file or a vault string.

### String Functions

- `GetString(key string) string`
- `GetStringOrDefault(key string, defaultValue string) string`
- `GetStringOrError(key string) (string, error)`
- `GetStringOrPanic(key string) string`

### Bool Functions

- `GetBool(key string) bool`
- `GetBoolOrDefault(key string, defaultValue bool) bool`
- `GetBoolOrError(key string) (bool, error)`
- `GetBoolOrPanic(key string) bool`

### Int Functions

- `GetInt(key string) int`
- `GetIntOrDefault(key string, defaultValue int) int`
- `GetIntOrError(key string) (int, error)`
- `GetIntOrPanic(key string) int`

### Float64 Functions

- `GetFloat64(key string) float64`
- `GetFloat64OrDefault(key string, defaultValue float64) float64`
- `GetFloat64OrError(key string) (float64, error)`
- `GetFloat64OrPanic(key string) float64`

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
API_TIMEOUT="30"
DEBUG_MODE="true"
```

### Basic Usage

```go
package main

import (
	"fmt"
	"github.com/dracory/env"
	"log"
)

func main() {
	// Load environment variables from .env files
	env.Load(".env")

	// Get a string variable with a default value
	dbHost := env.GetStringOrDefault("DB_HOST", "localhost")

	// Get a required string variable (panics if not set)
	secretKey := env.GetStringOrPanic("SECRET_KEY")

	// Get an integer variable, handle error if not found or invalid
	dbPort, err := env.GetIntOrError("DB_PORT")
	if err != nil {
		log.Fatalf("Invalid or missing DB_PORT: %v", err)
	}

	// Get a boolean variable
	debugMode := env.GetBool("DEBUG_MODE")

	fmt.Printf("Database: %s:%d
", dbHost, dbPort)
	fmt.Printf("Secret Key Loaded: %v
", secretKey != "")
	fmt.Printf("Debug Mode: %v
", debugMode)
}
```

### Value Processing
`GetString`, `GetStringOrDefault`, `GetStringOrError`, and `GetStringOrPanic` automatically process these prefixes:

- `base64:<encoded>` – Decodes using URL-safe base64.
- `obfuscated:<text>` – Deobfuscates using `github.com/gouniverse/envenc`.

This lets you safely store encoded/obfuscated values in `.env` or other sources while retrieving plain values at runtime.

### Boolean Parsing
Boolean functions (`GetBool`, etc.) parse values with flexibility:

- **True values**: `"true"`, `"True"`, `"TRUE"`, `"T"`, `"t"`, `"1"`, `"yes"`, `"Yes"`, `"YES"`, and any positive number.
- **False values**: `"false"`, `"False"`, `"FALSE"`, `"F"`, `"f"`, `"0"`, `"no"`, `"No"`, `"NO"`, and any negative number.
- Any other or empty value returns `false` (for `GetBool`) or the specified default.

### Advanced: Env Vault Loading
Env vault loading allows you to load environment variables from an encrypted vault file or a string.

```go
// Load environment variables from an encrypted vault file
err := env.LoadVault(struct {
    Password      string
    VaultFilePath string
    VaultContent  string
}{
    Password:      "your-password",
    VaultFilePath: ".env.vault",
})
if err != nil {
    // handle error
}
```

### Notes on Load
`Load()` will attempt to load from a default `.env` file, and then from any additional file paths you pass in. Missing files are silently skipped.


## Features

- Load environment variables from `.env` files (`Initialize`)
- Process values with `base64:` and `obfuscated:` prefixes automatically
- Simple and intuitive API
- Boolean helpers with sane defaults
- Optional encrypted vault loading (`VaultLoad`)

## API Reference

### Functions

- `Initialize(envFilePath ...string)` – Load environment variables from `.env` files. Defaults to `.env` and also attempts any additional paths provided.
- `Value(key string) string` – Get an environment variable. Returns empty string if not set. Automatically processes `base64:` and `obfuscated:` prefixes.
- `Must(key string) string` – Get a required environment variable. Panics if the variable is not set. Also processes prefixes.
- `Bool(key string) bool` – Get a boolean environment variable. Returns `false` if not set or invalid.
- `BoolDefault(key string, defaultValue bool) bool` – Get a boolean environment variable with a default value.

### Advanced Functions

- `VaultLoad(options struct{ Password string; VaultFilePath string; VaultContent string }) error` – Load environment variables from an encrypted vault file or a vault string. Exactly one of `VaultFilePath` or `VaultContent` must be provided along with a `Password`.

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
	
	// Get an environment variable (empty string if not set). Add your own default fallback if needed.
	dbHost := env.Value("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	
	// Get required environment variable (panics if not set)
	dbPort := env.Must("DB_PORT")
	
	fmt.Printf("Database: %s:%s\n", dbHost, dbPort)
}
```

### Value Processing
`Value` and `Must` automatically process these prefixes:

- `base64:<encoded>` – Decodes using URL-safe base64.
- `obfuscated:<text>` – Deobfuscates using `github.com/gouniverse/envenc`.

This lets you safely store encoded/obfuscated values in `.env` or other sources while retrieving plain values at runtime.

### Boolean Helpers
`Bool` and `BoolDefault` read booleans with flexible parsing:

- True values: `"true"`, `"True"`, `"TRUE"`, `"T"`, `"t"`, `"1"`, `"yes"`, `"Yes"`, `"YES"`
- False values: `"false"`, `"False"`, `"FALSE"`, `"F"`, `"f"`, `"0"`, `"no"`, `"No"`, `"NO"`
- Any other or empty value returns the provided default (or `false` for `Bool`).

### Advanced. Env Vault Loading
Env vault loading allows you to load environment variables from an encrypted vault file or a string.

```go
// Load environment variables from an encrypted vault file
err := env.VaultLoad(struct {
    Password      string
    VaultFilePath string
    VaultContent  string
}{
    Password:      "your-password",
    VaultFilePath: ".env.vault",
})
if err != nil {
    // handle error
}
```

### Notes on Initialize
`Initialize()` will attempt to load from a default `.env` file, and then from any additional file paths you pass in. Missing files are silently skipped.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is dual-licensed under the following terms:

- For non-commercial use, you may choose either the GNU Affero General Public License v3.0 (AGPLv3) or a separate commercial license (see below). You can find a copy of the AGPLv3 at: https://www.gnu.org/licenses/agpl-3.0.txt

- For commercial use, a separate commercial license is required. Commercial licenses are available for various use cases. Please contact me via my contact page at https://lesichkov.co.uk/contact to obtain a commercial license.
