package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()
	g.Use()

	f := func(c *gin.Context) {} // IRoutes cannot Group after Use
	g2 := g.GET("", f)           // embedded GET

	fmt.Println(g2 == g)

	gGrp := g.Group("/aa", f)
	gGrp.Use()

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
