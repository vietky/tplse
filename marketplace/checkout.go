package marketplace

import (
	"math"
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
	products     map[string]*ProductCount // key: productId int: count
	productRules []Rule                   // key: productId int: count
	globalRules  []Rule                   // key: productId int: count
}

// NewCheckout new checkout
func NewCheckout(rules []Rule) *Checkout {
	var productRules []Rule
	var globalRules []Rule
	for _, r := range rules {
		if r.Type == RuleTypeApplyAll {
			globalRules = append(globalRules, r)
		} else if r.Type == RuleTypeApplyToProduct {
			productRules = append(productRules, r)
		}
	}
	return &Checkout{
		productRules: productRules,
		globalRules:  globalRules,
		products:     make(map[string]*ProductCount),
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

func (c *Checkout) getTotalPriceForProductRule() (float64, error) {
	for _, rule := range c.productRules {
		for _, product := range c.products {
			// fmt.Printf("product %+v\n", product)
			if product.ProductCode == rule.Settings[ProductCode] {
				// try to parse all conditions
				v, err := strconv.ParseFloat(rule.Settings[PromotionPrice], 64)
				if err != nil {
					return 0, err
				}
				minCountCondition, err := strconv.Atoi(rule.Settings[MinCountCondition])
				if err != nil {
					return 0, err
				}
				// check condition
				if product.Count >= minCountCondition {
					product.Price = v
				}
			}
		}
	}
	result := 0.0
	for _, product := range c.products {
		result += product.Price * float64(product.Count)
	}
	return result, nil
}

// GetTotal calculate total price
func (c *Checkout) GetTotal() (float64, error) {
	result, err := c.getTotalPriceForProductRule()
	if err != nil {
		return result, err
	}
	for _, rule := range c.globalRules {
		// try to parse all conditions
		// fmt.Println("RuleTypeApplyAll", rule.Settings)
		discountPercent, err := strconv.Atoi(rule.Settings[DiscountPercent])
		if err != nil {
			return result, err
		}
		minPriceCondition, err := strconv.ParseFloat(rule.Settings[MinPriceCondition], 64)
		if err != nil {
			return 0, err
		}
		// check condition to get a discount
		if result >= (minPriceCondition) {
			result = result * float64(100-discountPercent) / 100
		}
	}
	// round up to 2 decimals
	return math.Round(result*100) / 100, nil
}
