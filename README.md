# Discard API Go SDK

Official Go SDK for the Discard API providing access to 500+ endpoints across multiple categories.

## DiscardAPI Go Client

[![Go Reference](https://pkg.go.dev/badge/github.com/stormfiber/discardapi-go.svg)](https://pkg.go.dev/github.com/stormfiber/discardapi-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/stormfiber/discardapi-go)](https://goreportcard.com/report/github.com/stormfiber/discardapi-go)
[![Test Status](https://github.com/stormfiber/discardapi-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/stormfiber/discardapi-go/actions/workflows/test.yml)
[![Lint Status](https://github.com/stormfiber/discardapi-go/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/stormfiber/discardapi-go/actions/workflows/lint.yml)
[![Codecov](https://codecov.io/gh/stormfiber/discardapi-go/branch/main/graph/badge.svg)](https://codecov.io/gh/stormfiber/discardapi-go)
![GitHub release](https://img.shields.io/github/v/release/stormfiber/discardapi-go)
![Go Version](https://img.shields.io/github/go-mod/go-version/stormfiber/discardapi-go)
[![License](https://img.shields.io/github/license/stormfiber/discardapi-go)](LICENSE)
![Workflow Status](https://img.shields.io/github/actions/workflow/status/stormfiber/discardapi-go/test.yml?branch=main)

---

## Overview

A Go client library for interacting with the **Discard API**.  
Provides a simple, typed interface for making API calls and handling responses efficiently.

## Installation

```bash
go get github.com/stormfiber/discardapi-go
```

## Quick Start
```go
package main

import (
    "fmt"
    "log"
    "time"
    
    discard "github.com/stormfiber/discardapi-go"
)

func main() {
    // Create client
    client, err := discard.NewClient(discard.Config{
        APIKey:       "your-api-key-here",
        FullResponse: false,
        Timeout:      30 * time.Second,
    })
    if err != nil {
        log.Fatal(err)
    }

    // Shorten URL
    result, err := client.ShortenClck("https://github.com/GlobalTechInfo")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Short URL:", result)

    // Get dad joke
    joke, err := client.DadJoke()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Joke:", joke)

    // Generate fake users
    users, err := client.FakeUsers(5, "en_US", "male", 1234)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Fake Users:", users)
}
```

## Features

✅ Full coverage of 500+ API endpoints

✅ Type-safe method signatures

✅ Configurable response modes (full or result-only)

✅ Built-in timeout handling

✅ Comprehensive error handling

✅ Zero external dependencies (stdlib only)

✅ Support for file uploads (FormData)

## Configuration
```bash
config := discard.Config{
    APIKey:       "your-api-key",   // Required
    BaseURL:      "https://...",     // Optional
    FullResponse: false,             // Optional
    Timeout:      30 * time.Second,  // Optional
}

client, err := discard.NewClient(config)
```

## API Categories

**Islamic**
Quran, Hadith, Prayer Times, Islamic Books

**AI**
Gemini Pro, Llama, Mistral, Image Generation

**Downloads**
Facebook, Instagram, TikTok, Twitter, YouTube, etc.

**URL Shortener**
Multiple shortening services

**Jokes & Quotes**
Dad jokes, programming jokes, motivational quotes

**Images**
Random images by category

**Image Makers**
QR codes, avatars, carbon images

**Faker**
Generate fake users, addresses, companies

**Music**
Spotify, SoundCloud, lyrics

**Facts**
Date facts, year facts, math facts

**Codec**
Base64, Base32, Binary encoding

## Examples

### Download TikTok Video
```go
video, err := client.DownloadTikTok("https://tiktok.com/@user/video/...")
if err != nil { log.Fatal(err) }
fmt.Println("Video URL:", video)
```
### Generate QR Code
```go
qr, err := client.QRTag("Hello World", "300x300", "255-0-0", "")
if err != nil {
    log.Fatal(err)
}
fmt.Println("QR Code:", qr)
```
### AI Text Generation
```go
response, err := client.GeminiPro("Explain quantum computing")
if err != nil {
    log.Fatal(err)
}
fmt.Println("AI Response:", response)
```
**Full Response Mode**
```go
client.SetFullResponse(true)
result, err := client.RandomQuote()
if err != nil {
    log.Fatal(err)
}
// result contains: {Creator, Result, Status}
fmt.Println(result)
```
**Error Handling**
```go
result, err := client.DownloadInstagram(url)
if err != nil {
    // Handle specific errors
    if strings.Contains(err.Error(), "timeout") {
        log.Println("Request timed out")
    } else if strings.Contains(err.Error(), "HTTP error") {
        log.Println("API returned error")
    } else {
        log.Println("Unknown error:", err)
    }
    return
}
```
## License
MIT © Qasim Ali
## Author
Qasim Ali
GitHub: @GlobalTechInfo
## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

**.github/workflows/lint.yml**

name: Lint

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  golangci:
    name: Lint
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - uses: actions/setup-go@v4
        with:
          go-version: '1.22'

      - name: Format code
        run: gofmt -s -w .

      - name: Run golangci-lint
        uses: golangci/golangci-lint-action@v3
        with:
          version: latest
          args: --out-format=colored-line-number
          
**.github/workflows/test.yml**

name: Test

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    strategy:
      matrix:
        go-version: [1.21.x, 1.22.x]
        os: [ubuntu-latest, macos-latest, windows-latest]
    runs-on: ${{ matrix.os }}

    steps:
      - uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: ${{ matrix.go-version }}

      - name: Install dependencies
        run: go mod download

      - name: Run short tests
        shell: bash
        run: go test -v -short -race -coverprofile=coverage.out -covermode=atomic ./...

      - name: Upload coverage
        if: matrix.os == 'ubuntu-latest' && matrix.go-version == '1.22.x'
        uses: codecov/codecov-action@v3
        with:
          files: ./coverage.out
          
