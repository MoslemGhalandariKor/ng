package accounts

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AddMembershipHandler(c *gin.Context) {
	membership := &Memberships{}

	membership.Email = c.PostForm("email")
	_, err := CreateMembership(membership)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

}


