package main

import (
	"github.com/gin-gonic/gin"
)

func updateHandler(c *gin.Context) {
	mhsReq, err := validatePostMhsHandler(c)
	if err != nil {
		failedResponseConstructor(c, err.Error())
		return
	}

	_, errCreate := createMhs(*mhsReq)
	if errCreate != nil {
		failedResponseConstructor(c, errCreate.Error())
		return
	}
	updateMhsSuccessResponseConstructor(c)
}
