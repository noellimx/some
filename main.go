package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type BaEngine struct {
	gin.IRoutes
	finisher finisher
}

func (b *BaEngine) finish(bS ...BaHandleFunc) gin.HandlerFunc {
	return someGinHandler
}
func (b *BaEngine) Use(handleFunc ...BaHandleFunc) BRoutes {
	r := b.IRoutes.Use(b.finish(handleFunc...))
	return &BaEngine{finisher: b.finisher, IRoutes: r}
}

func (b *BaEngine) Handle(s string, s2 string, handleFunc ...BaHandleFunc) BRoutes {
	r := b.IRoutes.Handle(s, s2, b.finish(handleFunc...))
	return &BaEngine{finisher: b.finisher, IRoutes: r}
}

func (b *BaEngine) Any(s string, handleFunc ...BaHandleFunc) BRoutes {
	r := b.IRoutes.Any(s, b.finish(handleFunc...))
	return &BaEngine{finisher: b.finisher, IRoutes: r}
}

func (b *BaEngine) GET(s string, handleFunc ...BaHandleFunc) BRoutes {
	r := b.IRoutes.GET(s, b.finish(handleFunc...))
	return &BaEngine{finisher: b.finisher, IRoutes: r}
}

func (b *BaEngine) POST(s string, handleFunc ...BaHandleFunc) BRoutes {
	r := b.IRoutes.POST(s, b.finish(handleFunc...))
	return &BaEngine{finisher: b.finisher, IRoutes: r}
}

func (b *BaEngine) DELETE(s string, handleFunc ...BaHandleFunc) BRoutes {
	r := b.IRoutes.POST(s, b.finish(handleFunc...))
	return &BaEngine{finisher: b.finisher, IRoutes: r}
}

func (b *BaEngine) PATCH(s string, handleFunc ...BaHandleFunc) BRoutes {
	r := b.IRoutes.PATCH(s, b.finish(handleFunc...))
	return &BaEngine{finisher: b.finisher, IRoutes: r}
}

func (b *BaEngine) PUT(s string, handleFunc ...BaHandleFunc) BRoutes {
	r := b.IRoutes.PUT(s, b.finish(handleFunc...))
	return &BaEngine{finisher: b.finisher, IRoutes: r}
}

func (b *BaEngine) OPTIONS(s string, handleFunc ...BaHandleFunc) BRoutes {
	r := b.IRoutes.OPTIONS(s, b.finish(handleFunc...))
	return &BaEngine{finisher: b.finisher, IRoutes: r}
}

func (b *BaEngine) HEAD(s string, handleFunc ...BaHandleFunc) BRoutes {
	r := b.IRoutes.HEAD(s, b.finish(handleFunc...))
	return &BaEngine{finisher: b.finisher, IRoutes: r}
}

func (b *BaEngine) Match(strings []string, s string, handleFunc ...BaHandleFunc) BRoutes {
	r := b.IRoutes.Match(strings, s, b.finish(handleFunc...))
	return &BaEngine{finisher: b.finisher, IRoutes: r}
}

func (b *BaEngine) StaticFile(s string, s2 string) BRoutes {
	b.IRoutes.StaticFile(s, s2)
	return b
}

func (b *BaEngine) StaticFileFS(s string, s2 string, system http.FileSystem) BRoutes {
	b.IRoutes.StaticFileFS(s, s2, system)
	return b
}

func (b *BaEngine) Static(s string, s2 string) BRoutes {
	b.IRoutes.Static(s, s2)
	return b
}

func (b *BaEngine) StaticFS(s string, system http.FileSystem) BRoutes {
	b.IRoutes.StaticFS(s, system)
	return b
}

type BGroup func(string, ...BaHandleFunc) gin.RouterGroup

type BRoutes interface {
	Use(...BaHandleFunc) BRoutes

	Handle(string, string, ...BaHandleFunc) BRoutes
	Any(string, ...BaHandleFunc) BRoutes
	GET(string, ...BaHandleFunc) BRoutes
	POST(string, ...BaHandleFunc) BRoutes
	DELETE(string, ...BaHandleFunc) BRoutes
	PATCH(string, ...BaHandleFunc) BRoutes
	PUT(string, ...BaHandleFunc) BRoutes
	OPTIONS(string, ...BaHandleFunc) BRoutes
	HEAD(string, ...BaHandleFunc) BRoutes
	Match([]string, string, ...BaHandleFunc) BRoutes

	StaticFile(string, string) BRoutes
	StaticFileFS(string, string, http.FileSystem) BRoutes
	Static(string, string) BRoutes
	StaticFS(string, http.FileSystem) BRoutes
}

type finisher func(*gin.Context) (int, interface{})

func defaultFinisher(*gin.Context) (int, interface{}) {
	return 1, struct{}{}
}
func NewBaEngine(engine *gin.Engine) *BaEngine {
	if engine != nil {
		log.Fatalf("init fail: nil engine")
	}

	return &BaEngine{
		IRoutes:  engine,
		finisher: defaultFinisher,
	}
}

// HandleFunc
func someMiddleWare(c *gin.Context) {
	c.Next()
}

func someGinHandler(c *gin.Context) {
	c.Next()
}

// BaHandleFunc
type BaHandleFunc finisher

var customMW BaHandleFunc = func(context *gin.Context) (int, interface{}) {
	return 1, nil
}

func main() {
	g := gin.New()
	g.Use(someMiddleWare)

	f := func(c *gin.Context) {}
	g2 := g.GET("", f) // embedded GET

	fmt.Println(g2 == g)

	gGrp := g.Group("/aa", f)
	gGrp.Use(someMiddleWare) // IRoutes cannot Group after Use
	gGrp.GET("/aa", func(c *gin.Context) {})
	h := NewBaEngine(g)

	h.Use(customMW)
	h.GET("/", func(ctx *gin.Context) (int, interface{}) {
		return 1, struct{}{}
	})
	g.HandleMethodNotAllowed = g.HandleMethodNotAllowed

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
