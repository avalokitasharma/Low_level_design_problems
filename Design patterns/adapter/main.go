package main

import (
	"errors"
	"fmt"
)

// Request type

type PaymentRequest struct {
	Amount     float64
	UPIID      string
	Account    string
	CardNumber string
	CVV        string
	IFSC       string
	Provider   string
}

// Common Interface

type PaymentProvider interface {
	Pay(req PaymentRequest) error
}

// PayTM gateway and adapter
type PayTMPaymentGateway struct{}

func (p *PayTMPaymentGateway) PayViaUPI(upi string, amount float64) {
	fmt.Printf("PayTM: Paid ₹%.2f via UPI %s\n", amount, upi)
}

// wraps over PayTM SDK

type PayTMAdapter struct {
	paytm *PayTMPaymentGateway
}

func NewPayTMAdapter() *PayTMAdapter {
	return &PayTMAdapter{paytm: &PayTMPaymentGateway{}}
}

// incoming payment request for PayTM
func (p *PayTMAdapter) Pay(req PaymentRequest) error {
	if req.UPIID == "" {
		return errors.New("PayTM requires a UPI ID")
	}
	p.paytm.PayViaUPI(req.UPIID, req.Amount)
	return nil
}

// Stripe SDK and adapter
type StripeSDK struct{}

func (s *StripeSDK) ChargeCard(card string, cvv string, cents int) {
	fmt.Printf("Stripe: Charged $%.2f to card %s\n", float64(cents)/100, card)
}

type StripeAdapter struct {
	stripe *StripeSDK
}

func NewStripeAdapter() *StripeAdapter {
	return &StripeAdapter{stripe: &StripeSDK{}}
}

func (a *StripeAdapter) Pay(req PaymentRequest) error {
	if req.CardNumber == "" || req.CVV == "" {
		return errors.New("stripe requires card number and CVV")
	}
	a.stripe.ChargeCard(req.CardNumber, req.CVV, int(req.Amount*100))
	return nil
}

// Razorpay client and adapter
type RazorpayClient struct{}

func (r *RazorpayClient) BankTransfer(account string, ifsc string, paise int) {
	fmt.Printf("Razorpay: Transferred ₹%.2f from A/C %s (%s)\n", float64(paise)/100, account, ifsc)
}

type RazorpayAdapter struct {
	razor *RazorpayClient
}

func NewRazorpayAdapter() *RazorpayAdapter {
	return &RazorpayAdapter{razor: &RazorpayClient{}}
}

func (a *RazorpayAdapter) Pay(req PaymentRequest) error {
	if req.Account == "" || req.IFSC == "" {
		return errors.New("razorpay requires account number and IFSC")
	}
	a.razor.BankTransfer(req.Account, req.IFSC, int(req.Amount*100))
	return nil
}

// common inteface/router
type PaymentProcessor struct {
	providers map[string]PaymentProvider
}

func NewPaymentProcessor() *PaymentProcessor {
	return &PaymentProcessor{
		providers: map[string]PaymentProvider{
			"paytm":    NewPayTMAdapter(),
			"stripe":   NewStripeAdapter(),
			"razorpay": NewRazorpayAdapter(),
		},
	}
}

func (pp *PaymentProcessor) Process(req PaymentRequest) {
	provider, ok := pp.providers[req.Provider]
	if !ok {
		fmt.Printf("Unknown payment provider: %s\n", req.Provider)
		return
	}
	err := provider.Pay(req)
	if err != nil {
		fmt.Println("Payment failed:", err)
	}
}

func main() {
	processor := NewPaymentProcessor()

	// PayTM request
	paytmReq := PaymentRequest{
		Amount:   150.75,
		UPIID:    "user@upi",
		Provider: "paytm",
	}
	processor.Process(paytmReq)

	// Stripe request
	stripeReq := PaymentRequest{
		Amount:     99.99,
		CardNumber: "4242-4242-4242-4242",
		CVV:        "123",
		Provider:   "stripe",
	}
	processor.Process(stripeReq)

	// Razorpay request
	razorReq := PaymentRequest{
		Amount:   500.00,
		Account:  "1234567890",
		IFSC:     "RAZR0000001",
		Provider: "razorpay",
	}
	processor.Process(razorReq)

	// missing info
	badReq := PaymentRequest{
		Amount:   100,
		Provider: "stripe",
	}
	processor.Process(badReq)
}
