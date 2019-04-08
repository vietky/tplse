package marketplace

// ProductRepository rule repo
type ProductRepository struct {
}

// GetAllProducts get promotional rules
func (*RuleRepository) GetAllProducts() []Product {
	rules := []Product{}
	rules = append(rules, Product{
		Code:  "001",
		Name:  "Lavender heart",
		Price: 9.25,
	})
	rules = append(rules, Product{
		Code:  "002",
		Name:  "Personalised cufflinks",
		Price: 45.0,
	})
	rules = append(rules, Product{
		Code:  "003",
		Name:  "Kids T-shirt ",
		Price: 19.95,
	})
	return rules
}

// GetProducts get products by id
func (repo *RuleRepository) GetProducts(ids []string) []Product {
	products := repo.GetAllProducts()
	result := []Product{}
	for _, code := range ids {
		for _, p := range products {
			if code == p.Code {
				result = append(result, p)
			}
		}
	}
	return result
}
