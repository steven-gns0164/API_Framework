package controllers

import (
	"github.com/gin-gonic/gin"
)

func sendSuccessResponseGIN(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message": message,
	})
}

func sendErrorResponseGIN(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message": message,
	})
}

// func sendSuccessResponse(w http.ResponseWriter, message string) {
// 	var response m.Response
// 	response.Status = 200
// 	response.Message = message
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func sendErrorResponse(w http.ResponseWriter, message string) {
// 	var response m.Response
// 	response.Status = 400
// 	response.Message = "Failed"
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }
