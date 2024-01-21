package bahttp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IRouter interface {
	IRoutes
	Group(string, ...HandlerBody) *RouterGroup
}

type IRoutes interface {
	Use(...gin.HandlerFunc) IRoutes

	Handle(string, string, ...HandlerBody) IRoutes
	Any(string, ...HandlerBody) IRoutes
	GET(string, ...HandlerBody) IRoutes
	POST(string, ...HandlerBody) IRoutes
	DELETE(string, ...HandlerBody) IRoutes
	PATCH(string, ...HandlerBody) IRoutes
	PUT(string, ...HandlerBody) IRoutes
	OPTIONS(string, ...HandlerBody) IRoutes
	HEAD(string, ...HandlerBody) IRoutes
	Match([]string, string, ...HandlerBody) IRoutes

	StaticFile(string, string) IRoutes
	StaticFileFS(string, string, http.FileSystem) IRoutes
	Static(string, string) IRoutes
	StaticFS(string, http.FileSystem) IRoutes
}

var _ IRouter = (*RouterGroup)(nil)

type RouterGroup struct {
	*gin.RouterGroup
	finisher Finisher
}

func (r *RouterGroup) Group(relativePath string, handlers ...HandlerBody) *RouterGroup {
	return &RouterGroup{
		RouterGroup: r.RouterGroup.Group(relativePath, r.unwrap(handlers...)...),
		finisher:    r.finisher,
	}
}

func (r *RouterGroup) Use(handlerFunc ...gin.HandlerFunc) IRoutes {
	ir := r.RouterGroup.Use(handlerFunc...)
	return r.returnObject(ir)
}

func (r *RouterGroup) Handle(s string, s2 string, b ...HandlerBody) IRoutes {
	ir := r.RouterGroup.Handle(s, s2, r.unwrap(b...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) Any(s string, b ...HandlerBody) IRoutes {
	ir := r.RouterGroup.Any(s, r.unwrap(b...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) GET(s string, b ...HandlerBody) IRoutes {
	ir := r.RouterGroup.GET(s, r.unwrap(b...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) POST(s string, b ...HandlerBody) IRoutes {
	ir := r.RouterGroup.POST(s, r.unwrap(b...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) DELETE(s string, b ...HandlerBody) IRoutes {
	ir := r.RouterGroup.DELETE(s, r.unwrap(b...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) PATCH(s string, b ...HandlerBody) IRoutes {
	ir := r.RouterGroup.PATCH(s, r.unwrap(b...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) PUT(s string, b ...HandlerBody) IRoutes {
	ir := r.RouterGroup.PUT(s, r.unwrap(b...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) OPTIONS(s string, b ...HandlerBody) IRoutes {
	ir := r.RouterGroup.OPTIONS(s, r.unwrap(b...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) HEAD(s string, b ...HandlerBody) IRoutes {
	ir := r.RouterGroup.HEAD(s, r.unwrap(b...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) Match(strings []string, s string, b ...HandlerBody) IRoutes {
	ir := r.RouterGroup.Match(strings, s, r.unwrap(b...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) StaticFile(s string, s2 string) IRoutes {
	ir := r.RouterGroup.StaticFile(s, s2)
	return r.returnObject(ir)
}

func (r *RouterGroup) StaticFileFS(s string, s2 string, system http.FileSystem) IRoutes {
	ir := r.RouterGroup.StaticFileFS(s, s2, system)
	return r.returnObject(ir)
}

func (r *RouterGroup) Static(s string, s2 string) IRoutes {
	ir := r.RouterGroup.Static(s, s2)
	return r.returnObject(ir)
}

func (r *RouterGroup) StaticFS(s string, system http.FileSystem) IRoutes {
	ir := r.RouterGroup.StaticFS(s, system)
	return r.returnObject(ir)
}

func (r *RouterGroup) unwrap(bS ...HandlerBody) gin.HandlersChain {
	chain := make(gin.HandlersChain, 0)

	for _, b := range bS {
		if b == nil {
			panic("nil bAhandler")
		}
		chain = append(chain, r.finisher(b))
	}

	return chain
}

func (r *RouterGroup) returnObject(ginIRoute gin.IRoutes) IRoutes {
	switch v := ginIRoute.(type) {

	case *gin.Engine:
		return &Engine{
			Engine:   v,
			finisher: r.finisher,
		}
	case *gin.RouterGroup:
		return &RouterGroup{
			RouterGroup: v,
			finisher:    r.finisher,
		}
	default:
		panic("unhandled concrete value of IRoutes")
	}

	return nil
}
