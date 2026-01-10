package routes

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func StoI(str string, defalt int) int {
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
