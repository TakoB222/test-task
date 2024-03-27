package service

import (
	"errors"
	"strconv"
	"time"
)

type CreditCardService struct{}

func NewCreditCardService() *CreditCardService {
	return &CreditCardService{}
}

func (c *CreditCardService) Validate(input CreditCardInput) (bool, error) {
	if err := validateExpiration(input); err != nil {
		return false, err
	}
	if validated := validateNumber(input.Number); !validated {
		return false, errors.New("invalid card number")
	}

	return false, nil
}

func validateExpiration(input CreditCardInput) (err error) {
	var year, month int
	now := time.Now()

	if len(input.Year) < 3 {
		year, err = strconv.Atoi(strconv.Itoa(now.UTC().Year())[:2] + input.Year)
		if err != nil {
			return errors.New("invalid year")
		}
	} else {
		year, err = strconv.Atoi(input.Year)
		if err != nil {
			return errors.New("invalid year")
		}
	}

	month, err = strconv.Atoi(input.Month)
	if err != nil {
		return errors.New("invalid month")
	}

	if month < 1 || 12 < month {
		return errors.New("invalid month")
	}

	if year < now.UTC().Year() {
		return errors.New("credit card has expired")
	}

	if year == now.UTC().Year() && month < int(now.UTC().Month()) {
		return errors.New("credit card has expired")
	}

	return nil
}

func validateNumber(number string) bool {
	var sum int
	var alternate bool

	numberLen := len(number)

	if numberLen < 13 || numberLen > 19 {
		return false
	}

	for i := numberLen - 1; i > -1; i-- {
		mod, _ := strconv.Atoi(string(number[i]))
		if alternate {
			mod *= 2
			if mod > 9 {
				mod = (mod % 10) + 1
			}
		}

		alternate = !alternate

		sum += mod
	}

	return sum%10 == 0
}
