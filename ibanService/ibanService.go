package main

import (
	"log"
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
	ibanValidator "github.com/mlenngren/iban-validator/ibanValidator"
)

type ValidateIbanRequest struct {
	Iban      string `json:"iban" binding:"required"`
	Validated int    `json:"validated"`
}

func main() {
	router := gin.Default()
	router.POST("/validate", func(c *gin.Context) {
		var p ValidateIbanRequest
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input!"})
			return
		}

		// Rudimentary ASCII check
		if !isAscii(p.Iban) {
			log.Printf("IBAN not in ASCII format")
			c.JSON(http.StatusBadRequest, p)
			return
		}

		p.Validated = ibanValidator.ValidateIban(p.Iban)
		log.Printf("ValidateIbanApiService: validated %v to %v", p.Iban, p.Validated)
		c.JSON(http.StatusOK, p)
	})
	router.Run(":5000")
}

func isAscii(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
