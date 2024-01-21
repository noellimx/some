package bahttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Finisher func(HandlerBody) gin.HandlerFunc
type HandlerBody func(c *gin.Context) int

type Engine struct {
	*gin.Engine
	finisher Finisher
}

var _ IRouter = (*Engine)(nil)

func (e *Engine) unwrap(bS ...HandlerBody) gin.HandlersChain {

	chain := make(gin.HandlersChain, 0)

	for _, b := range bS {
		if b == nil {
			panic("nil bAhandler")
		}

		chain = append(chain, e.finisher(b))
	}

	return chain
}

func (e *Engine) Group(relativePath string, handlers ...HandlerBody) *RouterGroup {
	return &RouterGroup{
		RouterGroup: e.Engine.Group(relativePath, e.unwrap(handlers...)...),
		finisher:    e.finisher,
	}
}

func (e *Engine) Use(b ...gin.HandlerFunc) IRoutes {
	ir := e.Engine.Use(b...)
	return e.returnObject(ir)
}

func (e *Engine) Handle(s string, s2 string, b ...HandlerBody) IRoutes {
	ir := e.Engine.Handle(s, s2, e.unwrap(b...)...)
	return e.returnObject(ir)
}

func (e *Engine) Any(s string, b ...HandlerBody) IRoutes {
	ir := e.Engine.Any(s, e.unwrap(b...)...)
	return e.returnObject(ir)
}

func (e *Engine) GET(s string, b ...HandlerBody) IRoutes {
	ir := e.Engine.GET(s, e.unwrap(b...)...)
	return e.returnObject(ir)
}

func (e *Engine) POST(s string, b ...HandlerBody) IRoutes {
	ir := e.Engine.POST(s, e.unwrap(b...)...)
	return e.returnObject(ir)
}

func (e *Engine) DELETE(s string, b ...HandlerBody) IRoutes {
	ir := e.Engine.DELETE(s, e.unwrap(b...)...)
	return e.returnObject(ir)
}

func (e *Engine) PATCH(s string, b ...HandlerBody) IRoutes {
	ir := e.Engine.PATCH(s, e.unwrap(b...)...)
	return e.returnObject(ir)
}

func (e *Engine) PUT(s string, b ...HandlerBody) IRoutes {
	ir := e.Engine.PUT(s, e.unwrap(b...)...)
	return e.returnObject(ir)
}

func (e *Engine) OPTIONS(s string, b ...HandlerBody) IRoutes {
	ir := e.Engine.OPTIONS(s, e.unwrap(b...)...)
	return e.returnObject(ir)
}

func (e *Engine) HEAD(s string, b ...HandlerBody) IRoutes {
	ir := e.Engine.HEAD(s, e.unwrap(b...)...)
	return e.returnObject(ir)
}

func (e *Engine) Match(strings []string, s string, b ...HandlerBody) IRoutes {
	ir := e.Engine.Match(strings, s, e.unwrap(b...)...)
	return e.returnObject(ir)
}

func (e *Engine) StaticFile(s string, s2 string) IRoutes {
	ir := e.Engine.StaticFile(s, s2)
	return e.returnObject(ir)
}

func (e *Engine) StaticFileFS(s string, s2 string, system http.FileSystem) IRoutes {
	ir := e.Engine.StaticFileFS(s, s2, system)
	return e.returnObject(ir)
}

func (e *Engine) Static(s string, s2 string) IRoutes {
	ir := e.Engine.Static(s, s2)
	return e.returnObject(ir)
}

func (e *Engine) StaticFS(s string, system http.FileSystem) IRoutes {
	ir := e.Engine.StaticFS(s, system)
	return e.returnObject(ir)
}

func (e *Engine) returnObject(ginIRoute gin.IRoutes) IRoutes {
	switch v := ginIRoute.(type) {

	case *gin.Engine:
		return &Engine{
			Engine:   v,
			finisher: e.finisher,
		}
	case *gin.RouterGroup:
		return &RouterGroup{
			RouterGroup: v,
			finisher:    e.finisher,
		}
	default:
		panic("unhandled concrete value of IRoutes")
	}

	return nil
}

func NewEngine() *Engine {
	return &Engine{
		Engine: gin.New(),
		finisher: func(body HandlerBody) gin.HandlerFunc {
			return func(c *gin.Context) {
				body(c)
			}
		},
	}
}
