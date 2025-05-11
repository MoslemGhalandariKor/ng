package sell

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"nextgen/internals/db/redis"
	"nextgen/internals/gintemplrenderer"
	"nextgen/pkg/product_management"
	"nextgen/templates/dashboard/pages/selling"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)
func AddToCart(c *gin.Context) {
    session := sessions.Default(c)
    username := session.Get("Username")
    if username == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
        return
    }

    productName := c.PostForm("product-name")
    if productName == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "product_name is required"})
        return
    }
	products, err := product_management.GetProductByNameService(productName)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get products from database"})
		return
	}

	productJSON, err := json.Marshal(products[0])

    if err != nil {

        log.Println("Failed to marshal cart item:", err)

        return

    }
    key := fmt.Sprintf("cart:user:%s", username)
    if err := redis.Client.RPush(redis.Ctx, key, productJSON).Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add to cart"})
        return
    }

    items, err := redis.Client.LRange(redis.Ctx, fmt.Sprintf("cart:%d", username), 0, -1).Result()

    for _, item := range items {

        var product product_management.ProductView

        err := json.Unmarshal([]byte(item), &product)

        if err != nil {

            log.Println("Failed to unmarshal cart item:", err)

            continue

        }

 

        products = append(products, product)

    }
    if err != nil {

        log.Println("Failed to get cart items:", err)

        

    }

    sellingPageProps := selling.SellingPageProps{
        Products: products,
    }

    // Respond with updated cart count for HTMX
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, selling.SellingPage(sellingPageProps))
	c.Render(http.StatusOK, r)
}

func ViewCart(c *gin.Context) {
    session := sessions.Default(c)
    username := session.Get("Username")
    if username == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
        return
    }

    key := fmt.Sprintf("cart:user:%s", username)
    items, err := redis.Client.SMembers(redis.Ctx, key).Result()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cart"})
        return
    }

    c.HTML(http.StatusOK, "partials/cart_view.templ", gin.H{"Items": items})
}

func RemoveFromCart(c *gin.Context) {
    session := sessions.Default(c)
    username := session.Get("Username")
    if username == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
        return
    }

    productID := c.PostForm("product_id")
    if productID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "product_id is required"})
        return
    }

    key := fmt.Sprintf("cart:user:%s", username)
    if err := redis.Client.SRem(redis.Ctx, key, productID).Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to remove item"})
        return
    }

    items, _ := redis.Client.SMembers(redis.Ctx, key).Result()

    // Respond with updated cart view (re-render entire cart)
    c.HTML(http.StatusOK, "partials/cart_view.templ", gin.H{"Items": items})
}
