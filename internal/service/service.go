package service

type (
	CreditCardInput struct {
		Number string
		Month  string
		Year   string
	}
)

type CreditCard interface {
	Validate(input CreditCardInput) (bool, error)
}

type Service struct {
	CreditCard
}

func NewServices() *Service {
	return &Service{
		CreditCard: NewCreditCardService(),
	}
}
