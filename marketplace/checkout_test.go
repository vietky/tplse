package marketplace

import (
	"strings"
	"testing"
)

func TestCheckout1(t *testing.T) {
	actual, err := checkout("001,002,003")
	if err != nil {
		t.Error(err)
	}
	if actual != 66.78 {
		t.Errorf("actual value %v is not equal to 66.78", actual)
	}
}

func TestCheckout2(t *testing.T) {
	actual, err := checkout("001,003,001")
	if err != nil {
		t.Error(err)
	}
	if actual != 36.95 {
		t.Errorf("actual value %v is not equal to 36.95", actual)
	}
}
func TestCheckout3(t *testing.T) {
	actual, err := checkout("001,002,001,003")
	if err != nil {
		t.Error(err)
	}
	if actual != 73.76 {
		t.Errorf("actual value %v is not equal to 73.76", actual)
	}
}

func checkout(input string) (float64, error) {
	ruleRepo := RuleRepository{}
	productRepo := ProductRepository{}

	promotionRules := ruleRepo.GetPromotionalRules()
	products := productRepo.GetProducts(strings.Split(input, ","))

	checkout := NewCheckout(promotionRules)
	for _, p := range products {
		checkout.Scan(&p)
	}

	return checkout.GetTotal()
}
