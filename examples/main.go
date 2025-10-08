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
		APIKey:       "YOUR_KEY_HERE",
		FullResponse: false,
		Timeout:      30 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🚀 Discard API Go SDK Examples")
	fmt.Println()

	// Example 1: URL Shortener
	fmt.Println("▶️  Example 1: URL Shortener")
	shortURL, err := client.ShortenClck("https://github.com/GlobalTechInfo")
	if err != nil {
		log.Printf("❌ Error: %v", err)
	} else {
		fmt.Printf("✅ Short URL: %v", shortURL)
	}
	fmt.Println()

	// Example 2: Dad Joke
	fmt.Println("▶️  Example 2: Dad Joke")
	joke, err := client.DadJoke()
	if err != nil {
		log.Printf("❌ Error: %v", err)
	} else {
		fmt.Printf("✅ Joke: %v", joke)
	}
	fmt.Println()

	// Example 3: Random Quote
	fmt.Println("▶️  Example 3: Random Quote")
	quote, err := client.RandomQuote()
	if err != nil {
		log.Printf("❌ Error: %v", err)
	} else {
		fmt.Printf("✅ Quote: %v", quote)
	}
	fmt.Println()

	// Example 4: Fake Users
	fmt.Println("▶️  Example 4: Fake Users")
	users, err := client.FakeUsers(3, "en_US", "female", 0)
	if err != nil {
		log.Printf("❌ Error: %v", err)
	} else {
		fmt.Printf("✅ Fake Users: %v", users)
	}
	fmt.Println()

	// Example 5: QR Code
	fmt.Println("▶️  Example 5: QR Code")
	qr, err := client.QRCode("Hello from Go SDK")
	if err != nil {
		log.Printf("❌ Error: %v", err)
	} else {
		fmt.Printf("✅ QR Code: %v", qr)
	}
	fmt.Println()

	// Example 6: Cat Image
	fmt.Println("▶️  Example 6: Cat Image")
	catImg, err := client.CatImage()
	if err != nil {
		log.Printf("❌ Error: %v", err)
	} else {
		fmt.Printf("✅ Cat Image: %v", catImg)
	}
	fmt.Println()

	// Example 7: Base64 Encoding
	fmt.Println("▶️  Example 7: Base64 Encoding")
	encoded, err := client.Base64("encode", "Hello World")
	if err != nil {
		log.Printf("❌ Error: %v", err)
	} else {
		fmt.Printf("✅ Encoded: %v", encoded)
	}
	fmt.Println()

	// Example 8: AI - Gemini Pro
	fmt.Println("▶️  Example 8: AI - Gemini Pro")
	aiResponse, err := client.GeminiPro("What is Go programming language?")
	if err != nil {
		log.Printf("❌ Error: %v", err)
	} else {
		fmt.Printf("✅ AI Response: %v", aiResponse)
	}
	fmt.Println()

	// Example 9: Date Fact
	fmt.Println("▶️  Example 9: Date Fact")
	dateFact, err := client.DateFact(7, 5)
	if err != nil {
		log.Printf("❌ Error: %v", err)
	} else {
		fmt.Printf("✅ Date Fact: %v", dateFact)
	}
	fmt.Println()

	// Example 10: Full Response Mode
	fmt.Println("▶️  Example 10: Full Response Mode")
	client.SetFullResponse(true)
	fullResp, err := client.RandomQuote()
	if err != nil {
		log.Printf("❌ Error: %v", err)
	} else {
		fmt.Printf("✅ Full Response: %v", fullResp)
	}
	fmt.Println()

	fmt.Println("✨ All examples completed!")
}
