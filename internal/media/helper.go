package media

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// ---------------------------- handler helpers ----------------------------
func isError(err error, errorMsg string, errorCode int, c *gin.Context) bool {
	if err != nil {
		err := fmt.Sprintf("%s: %s", errorMsg, err)
		log.Println(err)
		c.JSON(errorCode, gin.H{"error": err})
		return true
	}
	return false
}
