package marketplace

// Checkout checkout
type Checkout struct {
	promotionalRules []Rule
}

// Scan a product
func (*Checkout) Scan(p *Product) {

}
