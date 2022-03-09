package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mlenngren/iban-validator/ibanService/handlers"
)

// Setup the routes for our api, no middleware, auth, cors, just pure gin, for now
func Setup() *gin.Engine {
	r := gin.Default()

	r.POST("/validate", handlers.ValidateIban)

	return r
}
