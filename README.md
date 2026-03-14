# klogs-pgw-go

Klogs Payment Gateway official Go client library.

[![Go Version](https://img.shields.io/badge/go-1.25+-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

## Features

-  **Card Payment**: Direct card payments with 3D Secure support
-  **Provision & Commit**: Pre-authorization and capture workflows
-  **Refund & Void**: Transaction reversal operations
-  **Secure Authentication**: HMAC-SHA256 signature-based API authentication
-  **Hosted Payment Pages**: Integrated payment page solution
-  **Installment Options**: Query available installment options by BIN
-  **Payment Tokens**: Secure tokenization for card data

## Installation

```bash
go get github.com/klogs/klogs-pgw-go
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	klogs "github.com/klogs/klogs-pgw-go/client"
)

func main() {
	// Initialize client with your credentials
	client, err := klogs.New(
		"your-api-key",
		"your-secret-key",
		"https://pgw.klogs.dev",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create a payment
	resp, err := client.CardPayment.Pay(context.Background(), klogs.CreatePaymentRequest{
		Amount:        100.00,
		Installment:   1,
		ReferenceCode: "ORDER-2026-001",
		Card: &klogs.CreditCard{
			CardHolderName: "John Doe",
			CardNumber:     "5526080000000006",
			Cvv:            "123",
			ExpireMonth:    12,
			ExpireYear:     2027,
		},
		Currency:    "TRY",
		Email:       "customer@example.com",
		Phone:       "+905551234567",
		Use3d:       true,
		Explanation: "Product purchase",
	})

	if err != nil {
		log.Fatal(err)
	}

	if resp.Success {
		fmt.Printf("Payment successful! Behavior: %s\n", resp.Behavior)
		if resp.Link != "" {
			fmt.Printf("3D Secure redirect: %s\n", resp.Link)
		}
	} else {
		fmt.Printf("Payment failed: %s\n", resp.Error.Summary)
	}
}
```

## Usage Examples

### 1. Simple Card Payment (Non-3D)

```go
resp, err := client.CardPayment.Pay(ctx, klogs.CreatePaymentRequest{
	Amount:      50.00,
	Installment: 1,
	Card: &klogs.CreditCard{
		CardHolderName: "Jane Smith",
		CardNumber:     "4508034508034509",
		Cvv:            "000",
		ExpireMonth:    6,
		ExpireYear:     2028,
	},
	Currency: "TRY",
	Use3d:    false,
})
```

### 2. 3D Secure Payment

```go
resp, err := client.CardPayment.Pay(ctx, klogs.CreatePaymentRequest{
	Amount:      200.00,
	Installment: 3,
	Card: &klogs.CreditCard{
		CardHolderName: "Ali Yılmaz",
		CardNumber:     "5526080000000006",
		Cvv:            "123",
		ExpireMonth:    12,
		ExpireYear:     2027,
	},
	Currency:  "TRY",
	Use3d:     true,
	ReturnURL: "https://yoursite.com/payment/callback",
	Email:     "customer@example.com",
	Phone:     "+905551234567",
})

// If 3D Secure is required, redirect user to resp.Link
if resp.Behavior == "redirect" && resp.Link != "" {
	// Redirect user to 3D Secure page
	fmt.Println("Redirect to:", resp.Link)
}
```

### 3. Payment with Invoice & Shipping Address

```go
resp, err := client.CardPayment.Pay(ctx, klogs.CreatePaymentRequest{
	Amount:      150.00,
	Installment: 1,
	Card: &klogs.CreditCard{
		CardHolderName: "Mehmet Demir",
		CardNumber:     "5526080000000006",
		Cvv:            "456",
		ExpireMonth:    3,
		ExpireYear:     2029,
	},
	Currency: "TRY",
	Invoice: &klogs.Address{
		Name:        "Mehmet",
		Surname:     "Demir",
		Street1:     "Atatürk Caddesi No:123",
		City:        "Istanbul",
		District:    "Kadıköy",
		PostalCode:  "34710",
		CountryCode: "TR",
		Phone:       "+905551234567",
	},
	Shipping: &klogs.Address{
		Name:        "Mehmet",
		Surname:     "Demir",
		Street1:     "İş Merkezi Kat:5",
		City:        "Istanbul",
		District:    "Beşiktaş",
		PostalCode:  "34340",
		CountryCode: "TR",
		Phone:       "+905559876543",
	},
})
```

### 4. Provision Payment (Pre-authorization)

```go
// Create provision
resp, err := client.CardPayment.Pay(ctx, klogs.CreatePaymentRequest{
	Amount:      300.00,
	Installment: 1,
	Card: &klogs.CreditCard{
		CardHolderName: "Ayşe Kaya",
		CardNumber:     "5526080000000006",
		Cvv:            "789",
		ExpireMonth:    9,
		ExpireYear:     2028,
	},
	Currency:   "TRY",
	ChargeType: klogs.ChargeType_Provision,
})

// Later, commit the provision
commitResp, err := client.CardPayment.ProvisionCommit(ctx, klogs.ProvisionCommitRequest{
	ReferenceCode: "ORDER-2026-001",
	Amount:        300.00,
})
```

### 5. Query Installment Options

```go
installments, err := client.CardPayment.CommissionsByBinAsync(ctx, klogs.PaymentOptionQueryParams{
	Amount:    500.00,
	BinNumber: "552608",
	Currency:  "TRY",
})

if err == nil && installments.Success {
	fmt.Println("Available installment options retrieved")
}
```

### 6. Create Payment Token

```go
tokenResp, err := client.CardPayment.CreatePaymentToken(ctx)
if err == nil && tokenResp.Success {
	fmt.Printf("Token: %s\n", tokenResp.Token)
}
```

### 7. Hosted Payment Page

```go
hostedResp, err := client.HostedPayment.CreatePayment(ctx, klogs.HostedPaymentRequest{
	Amount:        250.00,
	Currency:      "TRY",
	ReferenceCode: "ORDER-2026-002",
	FullName:      "Fatma Öztürk",
	Email:         "fatma@example.com",
	Phone:         "+905551234567",
	ReturnURL:     "https://yoursite.com/payment/result",
	Explanation:   "Online shopping",
})

if hostedResp.Success {
	fmt.Printf("Payment ID: %s\n", hostedResp.PaymentId)
	fmt.Printf("Redirect to: %s\n", hostedResp.Link)
	// Redirect user to hostedResp.Link
}
```

## API Reference

### Client Initialization

```go
client, err := klogs.New(apiKey, secretKey, baseURL)
```

**Parameters:**
- `apiKey` (string): Your API key from Klogs dashboard
- `secretKey` (string): Your secret key for HMAC signature
- `baseURL` (string): API base URL (e.g., `https://pgw.klogs.dev`)

### CardPayment Methods

#### Pay
```go
Pay(ctx context.Context, model CreatePaymentRequest) (*CardPaymentResponse, error)
```
Create a new card payment with optional 3D Secure.

#### ProvisionCommit
```go
ProvisionCommit(ctx context.Context, model ProvisionCommitRequest) (*Response, error)
```
Commit a previously created provision (pre-authorization).

#### CreatePaymentToken
```go
CreatePaymentToken(ctx context.Context) (*PaymentTokenResponse, error)
```
Generate a payment token for secure transactions.

#### CommissionsByBinAsync
```go
CommissionsByBinAsync(ctx context.Context, params PaymentOptionQueryParams) (*PaymentTokenResponse, error)
```
Query available installment options by card BIN.

### HostedPayment Methods

#### CreatePayment
```go
CreatePayment(ctx context.Context, model HostedPaymentRequest) (*HostedPaymentResponse, error)
```
Create a hosted payment page session.

## Data Models

### CreatePaymentRequest
```go
type CreatePaymentRequest struct {
	Amount            float64
	Installment       int
	ReferenceCode     string
	Card              *CreditCard
	Currency          string
	Email             string
	Phone             string
	Use3d             bool
	ReturnURL         string
	ChargeType        ChargeType
	Invoice           *Address
	Shipping          *Address
	Explanation       string
	// ... additional fields
}
```

### CreditCard
```go
type CreditCard struct {
	CardHolderName string
	CardNumber     string
	Cvv            string
	ExpireMonth    int
	ExpireYear     int
}
```

### ChargeType Constants
```go
const (
	ChargeType_DirectSale ChargeType = "directSale"
	ChargeType_Provision  ChargeType = "provision"
)
```

## Error Handling

All API methods return an error as the second return value. Always check for errors:

```go
resp, err := client.CardPayment.Pay(ctx, request)
if err != nil {
	log.Printf("API error: %v", err)
	return
}

if !resp.Success {
	log.Printf("Payment failed: %s", resp.Error.Summary)
	return
}

// Process successful payment
```

## Security

- All API requests are authenticated using HMAC-SHA256 signatures
- Sensitive card data should be handled according to PCI DSS requirements
- Use 3D Secure (`Use3d: true`) for enhanced security
- Never log or store sensitive card information

## Testing

Use the test environment URL and credentials for development:

```go
client, _ := klogs.New(
	"test-api-key",
	"test-secret-key",
	"https://pgw.klogs.dev",
)
```

**Test Cards:**
- Visa: `4508034508034509`
- Mastercard: `5526080000000006`
- CVV: Any 3 digits
- Expiry: Any future date

## Support

- **Documentation**: [https://docs.klogs.io](https://docs.klogs.io)
- **Email**: support@klogs.io
- **Issues**: [GitHub Issues](https://github.com/klogs/klogs-pgw-go/issues)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

---

© 2025 Klogs. All rights reserved.