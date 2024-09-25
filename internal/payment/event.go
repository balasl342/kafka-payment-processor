package payment

// PaymentEvent represents the structure of a payment event
type PaymentEvent struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
}
