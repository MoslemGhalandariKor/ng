package product_management



func GetAllCategoriesService() (categories []CategoryView, err error) {

	categories, err = GetAllCategories()

	return categories, err

}

func GetAllProductsService() (products []ProductView, err error) {

	products, err = GetAllProducts()

	return products, err

}

func GetAllBrandsService() (brands []BrandView, err error) {

	brands, err = GetAllBrands()

	return brands, err

}