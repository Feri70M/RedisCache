# Redis Cache

A simple Redis caching library built with Go. This package provides a flexible and easy-to-use interface for caching data with support for both standalone and Sentinel Redis modes.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Methods](#methods)
- [Error Handling](#error-handling)

## Installation

To use this package in your project, you need to include it in your `go.mod` file. You can use the following command to add it:

```bash
go get github.com/~
```

## Usage

Here’s a simple example of how to use the Redis cache in your project:

### Example

```go
package main

import (
    "fmt"
    "log"
    "time"

    "get github.com/~"
)

func main() {

    // Configure the options for Redis
    opts := config.Options{
        Address:     "localhost:6379", // Change as needed
        Password:    "",
        DB:          0,
        UseSentinel: false, // Change to true if using Sentinel
    }

    // Create a new cache
    cache, err := config.NewCache(opts)
    if err != nil {
        log.Fatalf("Failed to create cache: %v", err)
    }
    defer cache.Close() // Ensure the client is closed at the end

    // Set a value in the cache
    err = cache.Set("key1", "value1", 10*time.Minute)
    if err != nil {
        log.Fatalf("Failed to set value: %v", err)
    }

    // Get a value from the cache
    value, err := cache.Get("key1")
    if err != nil {
        log.Fatalf("Failed to get value: %v", err)
    }
    fmt.Println("Retrieved value:", value)

    // Check if a key exists
    exists, err := cache.Exists("key1")
    if err != nil {
        log.Fatalf("Failed to check existence: %v", err)
    }
    fmt.Println("Does key1 exist?", exists)

    // Delete a key from the cache
    err = cache.Delete("key1")
    if err != nil {
        log.Fatalf("Failed to delete key: %v", err)
    }

    // Verify the key deletion
    exists, err = cache.Exists("key1")
    if err != nil {
        log.Fatalf("Failed to check existence: %v", err)
    }
    fmt.Println("Does key1 exist after deletion?", exists)
}

```

## Methods

### `Set(key string, val string, ttl time.Duration) error`
Sets the value for a key with an expiration time.

### `Get(key string) (string, error)`
Retrieves the value for a given key. Returns a `KeyNotFoundError` if the key does not exist.

### `Delete(key string) error`
Deletes a key from the cache.

### `Exists(key string) (bool, error)`
Checks if a key exists in the cache.

### `Close() error`
Closes the Redis client connection.


## Error Handling

The package defines a custom error type to enhance error handling:

- **KeyNotFoundError**: Returned when attempting to retrieve a non-existent key.

### KeyNotFoundError Implementation

```go
type KeyNotFoundError struct{}

func (e KeyNotFoundError) Error() string {
    return "key not found"
}