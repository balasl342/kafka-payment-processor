package fraud

import "github.com/balasl342/kafka-payment-processor/internal/payment"

const FraudThreshold = 10000.00

// IsFraudulent checks if the transaction amount exceeds the fraud threshold
func IsFraudulent(event payment.PaymentEvent) bool {
	return event.Amount > FraudThreshold
}
