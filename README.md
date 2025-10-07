# Discard API Go SDK

Official Go SDK for the Discard API providing access to 500+ endpoints across multiple categories.

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
