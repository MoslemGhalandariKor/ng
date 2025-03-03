package product_management

func GetAllProducts() ([]Product, error) {
	products, err := PrcGetAllProducts()
	return products, err
}
