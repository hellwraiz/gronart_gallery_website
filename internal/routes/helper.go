package routes

import (
	"log"
	"strconv"
) 

func StoI(str string, dflt int) int {
	out, err := strconv.Atoi(str)
	if err != nil { out = dflt; log.Panicf("Could not parse minPrice param: %s", err) }
	return out
}
