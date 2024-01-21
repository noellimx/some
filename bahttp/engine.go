package bahttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerBody func(c *gin.Context) (int, interface{})

type Finisher func(HandlerBody) gin.HandlerFunc

func defaultJSONFinisher() func(body HandlerBody) gin.HandlerFunc {
	return func(body HandlerBody) gin.HandlerFunc {
		return func(c *gin.Context) {
			httpStatus, resp := body(c)
			c.JSON(httpStatus, resp)
		}
	}
}

type Engine struct {
	*gin.Engine
	finisher Finisher
}

var _ IRouter = (*Engine)(nil)

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

func (e *Engine) Handle(httpMethod string, relativePath string, handlerBodies ...HandlerBody) IRoutes {
	ir := e.Engine.Handle(httpMethod, relativePath, e.unwrap(handlerBodies...)...)
	return e.returnObject(ir)
}

func (e *Engine) Any(httpMethod string, handlerBodies ...HandlerBody) IRoutes {
	ir := e.Engine.Any(httpMethod, e.unwrap(handlerBodies...)...)
	return e.returnObject(ir)
}

func (e *Engine) GET(relativePath string, handlerBodies ...HandlerBody) IRoutes {
	ir := e.Engine.GET(relativePath, e.unwrap(handlerBodies...)...)
	return e.returnObject(ir)
}

func (e *Engine) POST(relativePath string, handlerBodies ...HandlerBody) IRoutes {
	ir := e.Engine.POST(relativePath, e.unwrap(handlerBodies...)...)
	return e.returnObject(ir)
}

func (e *Engine) DELETE(relativePath string, handlerBodies ...HandlerBody) IRoutes {
	ir := e.Engine.DELETE(relativePath, e.unwrap(handlerBodies...)...)
	return e.returnObject(ir)
}

func (e *Engine) PATCH(relativePath string, handlerBodies ...HandlerBody) IRoutes {
	ir := e.Engine.PATCH(relativePath, e.unwrap(handlerBodies...)...)
	return e.returnObject(ir)
}

func (e *Engine) PUT(relativePath string, handlerBodies ...HandlerBody) IRoutes {
	ir := e.Engine.PUT(relativePath, e.unwrap(handlerBodies...)...)
	return e.returnObject(ir)
}

func (e *Engine) OPTIONS(relativePath string, handlerBodies ...HandlerBody) IRoutes {
	ir := e.Engine.OPTIONS(relativePath, e.unwrap(handlerBodies...)...)
	return e.returnObject(ir)
}

func (e *Engine) HEAD(relativePath string, handlerBodies ...HandlerBody) IRoutes {
	ir := e.Engine.HEAD(relativePath, e.unwrap(handlerBodies...)...)
	return e.returnObject(ir)
}

func (e *Engine) Match(httpMethods []string, relativePath string, handlerBodies ...HandlerBody) IRoutes {
	ir := e.Engine.Match(httpMethods, relativePath, e.unwrap(handlerBodies...)...)
	return e.returnObject(ir)
}

func (e *Engine) StaticFile(relativePath string, filepath string) IRoutes {
	ir := e.Engine.StaticFile(relativePath, filepath)
	return e.returnObject(ir)
}

func (e *Engine) StaticFileFS(relativePath string, filepath string, fs http.FileSystem) IRoutes {
	ir := e.Engine.StaticFileFS(relativePath, filepath, fs)
	return e.returnObject(ir)
}

func (e *Engine) Static(relativePath string, filepath string) IRoutes {
	ir := e.Engine.Static(relativePath, filepath)
	return e.returnObject(ir)
}

func (e *Engine) StaticFS(relativePath string, fs http.FileSystem) IRoutes {
	ir := e.Engine.StaticFS(relativePath, fs)
	return e.returnObject(ir)
}

func (e *Engine) SetFinisher(finisher Finisher) *Engine {
	return &Engine{
		Engine:   e.Engine,
		finisher: finisher,
	}
}

func (e *Engine) unwrap(handlerBodies ...HandlerBody) gin.HandlersChain {
	chain := make(gin.HandlersChain, 0)
	for _, b := range handlerBodies {
		if b == nil {
			panic("nil HandlerBody")
		}
		chain = append(chain, e.finisher(b))
	}
	return chain
}

func (e *Engine) returnObject(ginIRoutes gin.IRoutes) IRoutes {
	switch v := ginIRoutes.(type) {
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
}

func NewEngine() *Engine {
	return &Engine{
		Engine:   gin.New(),
		finisher: defaultJSONFinisher(),
	}
}
