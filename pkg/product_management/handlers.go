package product_management

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"
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
		Barcode:      c.PostForm("barcode"),
		CategoryID:   c.PostForm("category_id"),
		BrandID:      c.PostForm("brand_id"),
	}
	productsImagePath := "./static/images/product/"
	productImage, err := c.FormFile("image_file")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	imageExtension := filepath.Ext(productImage.Filename)

	newImageName := fmt.Sprintf("%d%s", time.Now().UnixNano(), imageExtension)

	imagePath := filepath.Join(productsImagePath, newImageName)
	if err := c.SaveUploadedFile(productImage, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	product.ImageSrc = imagePath

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
	c.Redirect(http.StatusFound, "/dashboard/Products")
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

func AddBrandHandler(c *gin.Context) {
	var brand Brand
	c.Bind(&brand)

	responseCode, responseDesc := AddBrand(brand)

	if responseCode != 0 {
		log.Panic(responseDesc)
	} else {
		c.Redirect(http.StatusFound, "/dashboard/brands")
	}

}

func DeleteBrandHandler(c *gin.Context) {
	idParam := c.Param("id")

	_, _ = DeleteBrandById(idParam)
	c.Redirect(http.StatusFound, "/dashboard/brands")
}

func AddWarehouseHandler(c *gin.Context) {
	var warehouse Warehouse
	c.Bind(&warehouse)

	responseCode, responseDesc := AddWarehouse(warehouse)

	if responseCode != 0 {
		log.Panic(responseDesc)
	} else {
		c.Redirect(http.StatusFound, "/dashboard/warehouses")
	}

}

func DeleteWarehouseHandler(c *gin.Context) {
	idParam := c.Param("id")

	_, _ = DeleteWarehouseById(idParam)
	c.Redirect(http.StatusFound, "/dashboard/warehouses")
}
