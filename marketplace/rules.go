package marketplace

// RuleRepository rule repo
type RuleRepository struct {
}

// GetPromotionalRules get promotional rules
func (*RuleRepository) GetPromotionalRules() []Rule {
	rules := []Rule{}
	rules = append(rules, Rule{
		Type: ApplyAll,
		Settings: map[RuleParam]string{
			PriceCondition:  "60.0",
			DiscountPercent: "10",
		},
	})
	rules = append(rules, Rule{
		Type: ApplyAll,
		Settings: map[RuleParam]string{
			ProductCode:    "001",
			PromotionPrice: "8.5",
			CountCondition: "2",
		},
	})
	return rules
}
