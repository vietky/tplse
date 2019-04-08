package marketplace

// RuleType define rule type
type RuleType int

const (
	// RuleTypeApplyAll to total price in checkout
	RuleTypeApplyAll RuleType = iota
	// RuleTypeApplyToProduct to a product
	RuleTypeApplyToProduct
)

// RuleParam define rule param names
type RuleParam string

const (
	ProductCode       RuleParam = "product_code"
	PromotionPrice    RuleParam = "promotion_price"
	MinCountCondition RuleParam = "min_count_condition"
	MinPriceCondition RuleParam = "min_price_condition"
	DiscountPercent   RuleParam = "discount_percent"
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

// ProductRule product rule
type ProductRule struct {
	ProductCode    string
	PromotionPrice float64
	CountCondition int
}

// CheckoutRule rule
type CheckoutRule struct {
	PriceCondition  float64
	DiscountPercent float64
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
	GetTotal()
}
