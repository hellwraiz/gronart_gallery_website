package paintings

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// --------------------------- repository helpers --------------------------
func generateUUID() string {
	return uuid.New().String()
}

// ---------------------------- handler helpers ----------------------------
func StoI(str string, defalt int) int {
	if str == "" {
		return defalt
	}
	out, err := strconv.Atoi(str)
	if err != nil {
		out = defalt
		log.Printf("Could not parse minPrice param: %s", err)
	}
	return out
}

func isError(err error, errorMsg string, errorCode int, c *gin.Context) bool {
	if err != nil {
		err := fmt.Sprintf("%s: %s", errorMsg, err)
		log.Println(err)
		c.JSON(errorCode, gin.H{"error": err})
		return true
	}
	return false
}
