package domain

type Offer struct {
	Message string
}

func NewOffer(message string) Offer {
	return Offer{Message: message}
}
