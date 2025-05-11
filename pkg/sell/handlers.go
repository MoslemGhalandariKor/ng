package sell


import (
    "fmt"
    "net/http"
    "nextgen/internals/db/redis"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
)

func AddToCart(c *gin.Context) {
    session := sessions.Default(c)
    username := session.Get("Username")
    if username == nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "not logged in"})
        return
    }

    productID := c.PostForm("product_id")
    if productID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "product_id required"})
        return
    }

    key := fmt.Sprintf("cart:user:%s", username.(string))
    err := redis.Client.SAdd(redis.Ctx, key, productID).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "redis error"})
        return
    }

    count, _ := redis.Client.SCard(redis.Ctx, key).Result()
    c.HTML(http.StatusOK, "partials/cart_count.templ", gin.H{"Count": count})
}
