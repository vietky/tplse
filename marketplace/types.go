package marketplace

// RuleType define rule type
type RuleType int

const (
	// ApplyAll to total price in checkout
	ApplyAll RuleType = iota
	// ApplyToProduct to a product
	ApplyToProduct
)

// RuleParam define rule param names
type RuleParam string

const (
	ProductCode     RuleParam = "product_code"
	PromotionPrice  RuleParam = "promotion_price"
	CountCondition  RuleParam = "count_condition"
	PriceCondition  RuleParam = "price_condition"
	PriceValue      RuleParam = "price_value"
	DiscountPercent RuleParam = "discount_percent"
)

// Product product info
type Product struct {
	Code  string
	Name  string
	Price float64
}

// Rule product rule
type Rule struct {
	// rule id is 404
	Type     RuleType
	Settings map[RuleParam]string
}

// IProductRepository product repo
type IProductRepository interface {
	GetAllProducts() []Product
	GetProducts([]string) []Product
}

// IRuleRepository rule repo
type IRuleRepository interface {
	GetPromotionalRules() []Rule
}

// ICheckout checkout
type ICheckout interface {
	Scan(Product)
}
