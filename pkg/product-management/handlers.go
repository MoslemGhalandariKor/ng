package product_management

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AddProductHandler(c *gin.Context) {
	product := Product{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		CategoryID:  "1",
	}

	responseCode, responseDesc := AddProduct(product)

	if responseCode != 0 {
		log.Panic(responseDesc)
	} else {
		c.Redirect(http.StatusFound, "/dashboard/products")
	}

}

func AddCategoryHandler(c *gin.Context) {
	category := Category{
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		ParentId:    c.PostForm("parent_id"),
	}

	responseCode, responseDesc := AddCategory(category)

	if responseCode != 0 {
		log.Panic(responseDesc)
	} else {
		c.Redirect(http.StatusFound, "/dashboard/categories")
	}

}
