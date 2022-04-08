package messages

import "github.com/google/uuid"

type BasketCheckoutMessage struct {
	UserId         uuid.UUID `json:"user_id"`
	CardNumber     string    `json:"card_number"`
	CardName       string    `json:"card_name"`
	CardExpiration string    `json:"card_expiration"`
	BasketTotal    float64   `json:"basket_total"`
}
