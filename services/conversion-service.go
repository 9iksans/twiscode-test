package services

import (
	// "fmt"
	"math"
	"strconv"
	"strings"
)

func BinaryToDecimal(binary []string) float64 {
	var value float64
	leng := len(binary)
	for _, v := range binary {
		leng = leng - 1
		if strings.Compare(v, "1") == 0 {
			value = value + (math.Pow(2, float64(leng)))
		}

	}
	return value
}

func DecimalToBinary(decimal int) string {
	var value []string

	for i := 7; i >= 0; i-- {
		value = append(value, strconv.Itoa(decimal%2))
		decimal = decimal / 2
	}
	return strings.Join(value, "")
}
