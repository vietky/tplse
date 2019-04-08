package marketplace

import (
	"fmt"
	"strings"
	"testing"
)

func TestCheckout1(t *testing.T) {
	checkout("001,002,003", t)
}

func TestCheckout2(t *testing.T) {
	checkout("001,003,001", t)
}
func TestCheckout3(t *testing.T) {
	checkout("001,002,001,003", t)
}

func checkout(input string, t *testing.T) {
	ruleRepo := RuleRepository{}
	productRepo := ProductRepository{}

	promotionRules := ruleRepo.GetPromotionalRules()
	products := productRepo.GetProducts(strings.Split(input, ","))

	checkout := NewCheckout(promotionRules)
	for _, p := range products {
		checkout.Scan(&p)
	}

	price, err := checkout.GetTotal()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(price)
}
