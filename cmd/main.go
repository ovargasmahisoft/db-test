package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ovargasmahisoft/db-test/dummy"
)

func main() {
	r := gin.Default()
	r.GET("v1/dummies", func(c *gin.Context) {
		d, err := dummy.FetchAllFromConnectionPool()

		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
		} else {
			c.JSON(200, gin.H{
				"items": d,
			})
		}

	})
	r.GET("v2/dummies", func(c *gin.Context) {
		d, err := dummy.FetchAllNewConnection()

		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
		} else {
			c.JSON(200, gin.H{
				"items": d,
			})
		}

	})
	r.Run(":5000")
}
