package main

import "github.com/gin-gonic/gin"

func failedResponseConstructor(c *gin.Context, err string) {
	c.JSON(400, gin.H{
		"error":             "Invalid Request",
		"error_description": err,
	})
}

func updateMhsSuccessResponseConstructor(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}
