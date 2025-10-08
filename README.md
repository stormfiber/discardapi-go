<h1 align="center"> Discard API Go SDK </h1>


<h2 align="center">
  Official Go SDK for the Discard API providing access to 500+ endpoints across multiple categories.
  A hub of RESTful APIs for developers, From downloaders and AI tools to image processing, games, and converters - everything you need to elevate your applications to new heights.
</h2>

[![Go Reference](https://pkg.go.dev/badge/github.com/stormfiber/discardapi-go.svg)](https://pkg.go.dev/github.com/stormfiber/discardapi-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/stormfiber/discardapi-go)](https://goreportcard.com/report/github.com/stormfiber/discardapi-go)
[![Test Status](https://github.com/stormfiber/discardapi-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/stormfiber/discardapi-go/actions/workflows/test.yml)
[![Lint Status](https://github.com/stormfiber/discardapi-go/actions/workflows/lint.yml/badge.svg?branch=main)](https://github.com/stormfiber/discardapi-go/actions/workflows/lint.yml)
![GitHub release](https://img.shields.io/github/v/release/stormfiber/discardapi-go)
![Go Version](https://img.shields.io/github/go-mod/go-version/stormfiber/discardapi-go)
[![License](https://img.shields.io/github/license/stormfiber/discardapi-go)](LICENSE)
![Workflow Status](https://img.shields.io/github/actions/workflow/status/stormfiber/discardapi-go/test.yml?branch=main)

---

## Overview

A Go client library for interacting with the **Discard API**.  
Provides a simple, typed interface for making API calls and handling responses efficiently.
Explore our complete catalog and documentation at [discardapi.dpdns.org](https://discardapi.dpdns.org)


## Getting Started with Discard Rest APIs

**Welcome to Discard Rest APIs, your one-stop solution for seamless API integrations! Our extensive collection of APIs is designed for developers building apps, businesses enhancing services, or tech enthusiasts experimenting with new ideas.**

**Step 1: Sign Up & Get Your API Key**
[Signup](https://discardapi.dpdns.org/auth)
- Create an account to access our API dashboard. Signing up is quick and easy, providing instant access to hundreds of powerful APIs.

**Step 2: Choose an API**
[Endpoints](https://discardapi.dpdns.org/)
- Browse our comprehensive API library and select the API that fits your needs. Each API includes detailed documentation with endpoints, parameters, and response formats.

**Step 3: Make Your First API Call**

- With your API key in hand, you're ready to start! All our APIs follow REST principles and are designed for simple, intuitive integration.

**Step 4: Some Features Are Premium Only 📊**

- For extensive usage and advanced features, upgrade to a PRO or VIP plan offering higher limits, faster response times, and premium feature access.

## Installation

Requirements

- Go 1.19 or later

- Internet connection (for API calls)

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
Base64, Base32, Binary encoding and **many more**

## Examples

### Download TikTok Video
```go
video, err := client.dlTikTok("https://tiktok.com/@user/video/...")
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
result, err := client.dlInstagram(url)
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

## Rate Limiting

The Discard API implements rate limiting to ensure fair usage across all users. Rate limit details and current quotas are available in your Discard API dashboard. The SDK does not implement automatic rate limit handling, so you should implement appropriate throttling mechanisms in your application code. Response headers may include rate limit information that you can access when using full response mode.

## Security Best Practices

When deploying applications using this SDK, follow security best practices to protect your API credentials and user data. Never commit API keys to version control systems or expose them in client-side code. Use environment variables to store sensitive configuration in production environments. Implement server-side proxies for browser applications to keep API keys secure. Rotate API keys periodically and immediately if you suspect compromise. Validate and sanitize all user input before passing it to SDK methods.

## Contributing

Contributions to the Discard API SDK are welcome and appreciated. To contribute, begin by forking the repository at https://github.com/stormfiber/discardapi-go. Create a feature branch for your changes using descriptive naming conventions. Make your modifications while ensuring existing tests continue to pass. Add new tests to cover your changes when applicable. Submit a pull request with a clear description of your improvements.

## License

This SDK is released under the MIT License, granting permission to use, copy, modify, merge, publish, distribute, sublicense, and sell copies of the software. The software is provided without warranty of any kind, express or implied. The complete license text is available in the LICENSE file included with the SDK distribution.

## Support and Resources

For comprehensive API documentation and endpoint references, visit the official Discard API documentation at https://discardapi.dpdns.org. Report bugs and request features through the GitHub issue tracker at https://github.com/stormfiber/discardapi-go/issues. View the complete source code and contribute to development at https://github.com/stormfiber/discardapi-go. For general questions and community discussion, join our Discord community or open a GitHub discussion.

## Changelog

Version 1.0.0 represents the initial release of the Discard API SDK.

## Author

This SDK was created and is maintained by Qasim Ali. Connect with the author on GitHub at https://github.com/GlobalTechInfo or visit the project homepage for updates and announcements.

## Acknowledgments

Special thanks to all contributors who have helped improve this SDK through bug reports, feature requests, and code contributions. This SDK builds upon the robust Discard API service, and we appreciate the ongoing development and maintenance of the underlying API infrastructure. The open source community provides invaluable tools and libraries that make projects like this possible.

## Made with ❤️ by Qasim Ali

Package Name: discardapi-go

Version: 1.0.0

Repository: https://github.com/stormfiber/discardapi-go

## 🎧 Support & Contact

Need help or want to upgrade? Our team is here to assist you with integration, troubleshooting, and custom solutions.

### 📞 Contact Methods

- ![](https://img.shields.io/badge/Email-blue) **Email Support** → [discardapi@gmail.com](mailto:discardapi@gmail.com?subject=Support&body=Hello%20Team,)  
- ![](https://img.shields.io/badge/WhatsApp-green) **Live Chat** → [Chat on WhatsApp](https://wa.me/923051391007?text=Hello%20I%20need%20support)  
- ![](https://img.shields.io/badge/Discord-purple) **Community Support** → [Join our Discord Server](https://discord.gg/YBkzCWqz)  
- ![](https://img.shields.io/badge/GitHub-black) **Documentation** → [GitHub Examples](https://github.com/GlobalTechInfo/discard-api)  

---

### ⏱ Response Times
- ✅ **Premium Support** → < 1 hour  
- 🆓 **Free Support** → < 24 hours  

---
© 2025 **Discard API** — Built with Go & Fiber
          
