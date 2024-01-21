package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"some/bahttp"
)

func main() {
	r2 := bahttp.NewEngine()

	r2.Group("/aa").Group("/bb").Any("/hello", func(c *gin.Context) (int, interface{}) {
		return -2, struct{}{}
	})

	r2.Group("/cc").SetFinisher(func(b bahttp.HandlerBody) gin.HandlerFunc {
		return func(c *gin.Context) {
			a, _ := b(c)
			fmt.Printf("okie %d \n", a)
		}
	}).Any("/hello", func(c *gin.Context) (int, interface{}) {
		return -2, struct{}{}
	})
	_ = r2.Run(":8081")
}
