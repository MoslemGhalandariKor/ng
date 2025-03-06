package product_management

import (
	"log"
	"net/http"

	// "strconv"
	"github.com/gin-gonic/gin"
)

func AddProductHandler(c *gin.Context) {
	product := Product{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		Price:       c.PostForm("price"),
		// Amount:		   c.PostForm("amount"),
		ProdSize:     c.PostForm("prod_size"),
		ProdLength:   c.PostForm("prod_length"),
		ProdMaterial: c.PostForm("prod_material"),
		ProdColor:    c.PostForm("prod_color"),
		ImageSrc:     c.PostForm("image_src"),
		Barcode:      c.PostForm("barcode"),
		CategoryID:   c.PostForm("category_id"),
		BrandID:      c.PostForm("brand_id"),
	}
	responseCode, responseDesc := AddProduct(product)

	if responseCode != 0 {
		log.Panic(responseDesc)
	} else {
		c.Redirect(http.StatusFound, "/dashboard/products")
	}

}


func DeleteProductHandler(c *gin.Context) {
	idParam := c.Param("id")

	_, _ = DeleteProductById(idParam)
	c.Redirect(http.StatusFound, "/dashboard/Product")
}


func AddCategoryHandler(c *gin.Context) {
	var category Category
	c.Bind(&category)

	responseCode, responseDesc := AddCategory(category)

	if responseCode != 0 {
		log.Panic(responseDesc)
	} else {
		c.Redirect(http.StatusFound, "/dashboard/categories")
	}

}
func DeleteCategoryHandler(c *gin.Context) {
	idParam := c.Param("id")

	_, _ = DeleteCategoryById(idParam)
	c.Redirect(http.StatusFound, "/dashboard/categories")
}
