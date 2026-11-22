package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(c *gin.Context, data interface{}) bool {
	if err := c.ShouldBindJSON(data); err != nil {

		if ve, ok := err.(validator.ValidationErrors); ok {
			out := make(map[string][]string)
			for _, fe := range ve {

				out[fe.Field()] = append(out[fe.Field()], "The "+fe.Field()+" field is invalid based on "+fe.Tag())
			}
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"success": false,
				"message": "Validation Error",
				"data":    out,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Invalid JSON format",
			})
		}
		return false
	}
	return true
}
