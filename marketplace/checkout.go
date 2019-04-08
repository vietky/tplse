package marketplace

import (
	"strconv"
)

// ProductCount product count
type ProductCount struct {
	ProductCode string
	Price       float64
	Count       int
}

// Checkout checkout
type Checkout struct {
	promotionalRules []Rule
	products         map[string]*ProductCount // key: productId int: count
}

// NewCheckout new checkout
func NewCheckout(rules []Rule) *Checkout {
	return &Checkout{
		promotionalRules: rules,
		products:         make(map[string]*ProductCount),
	}
}

// Scan a product
func (c *Checkout) Scan(p *Product) {
	if _, ok := c.products[p.Code]; !ok {
		c.products[p.Code] = &ProductCount{
			ProductCode: p.Code,
			Price:       p.Price,
			Count:       0,
		}
	}
	c.products[p.Code].Count++
}

// GetTotal calculate total price
func (c *Checkout) GetTotal() (float64, error) {
	for _, rule := range c.promotionalRules {
		if rule.Type == RuleTypeApplyAll {
			continue
		}
		for _, product := range c.products {
			// fmt.Printf("product %+v\n", product)
			if product.ProductCode == rule.Settings[ProductCode] {
				v, err := strconv.ParseFloat(rule.Settings[PromotionPrice], 64)
				if err != nil {
					return 0, err
				}
				minCountCondition, err := strconv.Atoi(rule.Settings[MinCountCondition])
				if err != nil {
					return 0, err
				}
				if product.Count > minCountCondition {
					product.Price = v
				}
			}
		}
	}
	result := 0.0
	for _, product := range c.products {
		result += product.Price * float64(product.Count)
	}
	for _, rule := range c.promotionalRules {
		// fmt.Println("RuleTypeApplyAll", rule.Settings)
		if rule.Type == RuleTypeApplyToProduct {
			continue
		}
		discountPercent, err := strconv.Atoi(rule.Settings[DiscountPercent])
		if err != nil {
			return result, err
		}
		minPriceCondition, err := strconv.ParseFloat(rule.Settings[MinPriceCondition], 64)
		if err != nil {
			return 0, err
		}
		if result > (minPriceCondition) {
			result = result * float64(100-discountPercent) / 100
		}
	}
	return result, nil
}
