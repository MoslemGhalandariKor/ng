package product_management

func GetAllCategoriesService() (categories []CategoryView, err error) {

	categories, err = GetAllCategories()

	return categories, err

}

func GetCategoriesByPatternService(pattern string) (categories []CategoryView, err error) {

	categories, err = GetCategoriesByPattern(pattern)

	return categories, err

}

func GetAllProductsService() (products []ProductView, err error) {

	products, err = GetAllProducts()

	return products, err

}
func GetProductByNameService(productName string) (products []ProductView, err error) {

	products, err = GetProductByName(productName)

	return products, err

}

func GetAllBrandsService() (brands []BrandView, err error) {

	brands, err = GetAllBrands()

	return brands, err

}

func GetBrandsByPatternService(pattern string) (brands []BrandView, err error) {

	brands, err = GetBrandsByPattern(pattern)

	return brands, err

}
