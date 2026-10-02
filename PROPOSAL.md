# Proposal: Array Element Transformations and Case Normalization in `env`

## Overview

This document proposes design options for handling array element case normalization and custom item transformations when retrieving array environment variables using `github.com/dracory/env`.

---

## 1. Problem Statement

When retrieving array environment variables (such as allowed login methods `LOGIN_METHODS="password, WebAuthn, OAuth"` or permissions/scopes), applications frequently require case normalization (typically lowercasing or uppercasing) or custom trimming/formatting.

Currently, developers must write repetitive boilerplate after calling `env.GetArray(...)`:

```go
// Current workaround
rawMethods := env.GetArray("LOGIN_METHODS")
loginMethods := make([]string, 0, len(rawMethods))
for _, m := range rawMethods {
    loginMethods = append(loginMethods, strings.ToLower(m))
}
```

Integrating case normalization and mapping directly into the `env` package reduces boilerplate, prevents bugs, and keeps environment configuration code concise.

---

## 2. Proposed Options

### Option 1: Dedicated Case Normalization Helpers (`GetArrayLower` / `GetArrayUpper`)

Add dedicated convenience functions specifically for lowercasing and uppercasing array elements.

#### Proposed API
```go
// Lowercase variants
func GetArrayLower(key string, separators ...string) []string
func GetArrayLowerOrDefault(key string, defaultValue []string, separators ...string) []string
func GetArrayLowerOrError(key string, separators ...string) ([]string, error)
func GetArrayLowerOrPanic(key string, separators ...string) []string

// Uppercase variants
func GetArrayUpper(key string, separators ...string) []string
func GetArrayUpperOrDefault(key string, defaultValue []string, separators ...string) []string
func GetArrayUpperOrError(key string, separators ...string) ([]string, error)
func GetArrayUpperOrPanic(key string, separators ...string) []string
```

#### Usage Example
```go
loginMethods := env.GetArrayLower("LOGIN_METHODS")
// Input: "Password, WebAuthn, OAUTH"
// Output: []string{"password", "webauthn", "oauth"}
```

#### Pros
- **Highly Ergonomic & Intuitive:** Solves the 90%+ use case (lowercasing/uppercasing) with zero boilerplate.
- **Consistent with `env` API Paradigm:** Fits seamlessly alongside `GetString`, `GetBool`, `GetInt`, `GetFloat`.
- **Zero Function Closure Overhead:** Clean call sites without needing `func(s string) string`.

#### Cons
- Expands package API surface by adding 8 new top-level functions (and matching `Validator` methods).

---

### Option 2: Higher-Order Mapping Functions (`GetArrayMapped` / `GetArrayMap`)

Add generic mapping functions that accept a transformation function `mapper func(string) string`.

#### Proposed API
```go
func GetArrayMapped(key string, mapper func(string) string, separators ...string) []string
func GetArrayMappedOrDefault(key string, defaultValue []string, mapper func(string) string, separators ...string) []string
func GetArrayMappedOrError(key string, mapper func(string) string, separators ...string) ([]string, error)
func GetArrayMappedOrPanic(key string, mapper func(string) string, separators ...string) []string
```

#### Usage Example
```go
loginMethods := env.GetArrayMapped("LOGIN_METHODS", strings.ToLower)
// Output: []string{"password", "webauthn", "oauth"}

// Custom transformation
cleanTags := env.GetArrayMapped("TAGS", func(s string) string {
    return strings.ReplaceAll(strings.ToLower(s), "-", "_")
})
```

#### Pros
- **Maximum Flexibility:** Supports lowercasing, uppercasing, sanitization, regex replacement, prefix/suffix stripping, etc.
- **Smaller API Addition:** Requires only 4 functions (`Mapped`, `MappedOrDefault`, `MappedOrError`, `MappedOrPanic`).

#### Cons
- Slightly more verbose for the common case (`env.GetArrayMapped("KEY", strings.ToLower)` vs `env.GetArrayLower("KEY")`).

---

### Option 3: Hybrid Approach (Recommended)

Combine Option 1 and Option 2 by providing `GetArrayLower` / `GetArrayUpper` for the most common use cases, backed by `GetArrayMapped` under the hood.

#### Implementation Architecture
```go
// GetArrayMapped applies custom mapper transformation to each non-empty element
func GetArrayMapped(key string, mapper func(string) string, separators ...string) []string {
    elems := GetArray(key, separators...)
    if elems == nil || mapper == nil {
        return elems
    }
    res := make([]string, len(elems))
    for i, e := range elems {
        res[i] = mapper(e)
    }
    return res
}

// GetArrayLower convenience wrapper
func GetArrayLower(key string, separators ...string) []string {
    return GetArrayMapped(key, strings.ToLower, separators...)
}

// GetArrayUpper convenience wrapper
func GetArrayUpper(key string, separators ...string) []string {
    return GetArrayMapped(key, strings.ToUpper, separators...)
}
```

---

## 3. Comparison Matrix

| Feature / Criteria | Option 1 (`GetArrayLower`/`Upper`) | Option 2 (`GetArrayMapped`) | Option 3 (Hybrid) |
| :--- | :--- | :--- | :--- |
| **Ease of Use (Common Case)** | ⭐⭐⭐⭐⭐ (Minimal typing) | ⭐⭐⭐⭐ (Requires passing `strings.ToLower`) | ⭐⭐⭐⭐⭐ (Best of both worlds) |
| **Flexibility** | ⭐⭐ (Only lower/upper) | ⭐⭐⭐⭐⭐ (Any `func(string) string`) | ⭐⭐⭐⭐⭐ (Full flexibility + shortcuts) |
| **API Consistency** | ⭐⭐⭐⭐⭐ (Matches `env` style) | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| **API Surface Impact** | +8 functions | +4 functions | +12 functions (or +6 core functions) |

---

## 4. Recommendation

We recommend **Option 3 (Hybrid Approach)**:

1. Provide `GetArrayMapped`, `GetArrayMappedOrDefault`, `GetArrayMappedOrError`, `GetArrayMappedOrPanic` as the general transformation foundation.
2. Provide `GetArrayLower` and `GetArrayUpper` (and their `OrDefault`/`OrError`/`OrPanic` variants) as convenience methods built on top of `GetArrayMapped`.
3. Add corresponding methods to `Validator` for consistency.

This provides the ultimate developer experience: single-function call convenience for standard lowercasing/uppercasing, plus complete power and flexibility for arbitrary string transformations.
