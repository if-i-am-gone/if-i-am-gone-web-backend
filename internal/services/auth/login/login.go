package login

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ididie/ifidie_backend/internal/database/pg"
)

func Login(c *gin.Context) {

	user, err := pg.PgClient.GetUser(c, 1)

	if err != nil {
		fmt.Printf(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	fmt.Printf("user found: %s", user.FirstName)

	c.JSON(http.StatusOK, gin.H{"data": user})

}
