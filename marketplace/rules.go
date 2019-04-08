package marketplace

// RuleRepository rule repo
type RuleRepository struct {
}

// GetPromotionalRules get promotional rules
func (*RuleRepository) GetPromotionalRules() []Rule {
	rules := []Rule{}
	rules = append(rules, Rule{
		Type: RuleTypeApplyAll,
		Settings: map[RuleParam]string{
			MinPriceCondition: "60.0",
			DiscountPercent:   "10",
		},
	})
	rules = append(rules, Rule{
		Type: RuleTypeApplyToProduct,
		Settings: map[RuleParam]string{
			ProductCode:       "001",
			MinCountCondition: "2",
			PromotionPrice:    "8.5",
		},
	})
	return rules
}
