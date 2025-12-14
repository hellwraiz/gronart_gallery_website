package routes

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StoI(str string, defalt int) int {
	out, err := strconv.Atoi(str)
	if err != nil { out = defalt; log.Panicf("Could not parse minPrice param: %s", err) }
	return out
}

func isError(err error, errorMsg string, errorCode int, c *gin.Context) bool {
	if err != nil {
		err := fmt.Sprintf("%s: %s", errorMsg, err)
		log.Printf(err)
		c.JSON(errorCode, gin.H{"error": err})
		return true
	}
	return false
}
