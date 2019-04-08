package marketplace

import "strings"

// Run parse input string and start checking out
func Run(input string) (float64, error) {
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
