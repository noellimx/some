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

type RouterGroup struct {
	*gin.RouterGroup
	finisher Finisher
}

var _ IRouter = (*RouterGroup)(nil)

func (r *RouterGroup) Group(relativePath string, handlerBodies ...HandlerBody) *RouterGroup {
	return &RouterGroup{
		RouterGroup: r.RouterGroup.Group(relativePath, r.unwrap(handlerBodies...)...),
		finisher:    r.finisher,
	}
}

func (r *RouterGroup) Use(handlerFunc ...gin.HandlerFunc) IRoutes {
	ir := r.RouterGroup.Use(handlerFunc...)
	return r.returnObject(ir)
}

func (r *RouterGroup) Handle(httpMethod string, s2 string, handlerBodies ...HandlerBody) IRoutes {
	ir := r.RouterGroup.Handle(httpMethod, s2, r.unwrap(handlerBodies...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) Any(httpMethod string, handlerBodies ...HandlerBody) IRoutes {
	ir := r.RouterGroup.Any(httpMethod, r.unwrap(handlerBodies...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) GET(httpMethod string, handlerBodies ...HandlerBody) IRoutes {
	ir := r.RouterGroup.GET(httpMethod, r.unwrap(handlerBodies...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) POST(httpMethod string, handlerBodies ...HandlerBody) IRoutes {
	ir := r.RouterGroup.POST(httpMethod, r.unwrap(handlerBodies...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) DELETE(httpMethod string, handlerBodies ...HandlerBody) IRoutes {
	ir := r.RouterGroup.DELETE(httpMethod, r.unwrap(handlerBodies...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) PATCH(httpMethod string, handlerBodies ...HandlerBody) IRoutes {
	ir := r.RouterGroup.PATCH(httpMethod, r.unwrap(handlerBodies...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) PUT(httpMethod string, handlerBodies ...HandlerBody) IRoutes {
	ir := r.RouterGroup.PUT(httpMethod, r.unwrap(handlerBodies...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) OPTIONS(httpMethod string, handlerBodies ...HandlerBody) IRoutes {
	ir := r.RouterGroup.OPTIONS(httpMethod, r.unwrap(handlerBodies...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) HEAD(httpMethod string, handlerBodies ...HandlerBody) IRoutes {
	ir := r.RouterGroup.HEAD(httpMethod, r.unwrap(handlerBodies...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) Match(strings []string, httpMethod string, handlerBodies ...HandlerBody) IRoutes {
	ir := r.RouterGroup.Match(strings, httpMethod, r.unwrap(handlerBodies...)...)
	return r.returnObject(ir)
}

func (r *RouterGroup) StaticFile(relativePath string, filepath string) IRoutes {
	ir := r.RouterGroup.StaticFile(relativePath, filepath)
	return r.returnObject(ir)
}

func (r *RouterGroup) StaticFileFS(relativePath string, filepath string, fs http.FileSystem) IRoutes {
	ir := r.RouterGroup.StaticFileFS(relativePath, filepath, fs)
	return r.returnObject(ir)
}

func (r *RouterGroup) Static(relativePath string, filepath string) IRoutes {
	ir := r.RouterGroup.Static(relativePath, filepath)
	return r.returnObject(ir)
}

func (r *RouterGroup) StaticFS(relativePath string, fs http.FileSystem) IRoutes {
	ir := r.RouterGroup.StaticFS(relativePath, fs)
	return r.returnObject(ir)
}

func (r *RouterGroup) unwrap(bS ...HandlerBody) gin.HandlersChain {
	chain := make(gin.HandlersChain, 0)

	for _, b := range bS {
		if b == nil {
			panic("nil HandlerBody")
		}
		chain = append(chain, r.finisher(b))
	}

	return chain
}

func (r *RouterGroup) SetFinisher(finisher Finisher) *RouterGroup {
	return &RouterGroup{
		RouterGroup: r.RouterGroup,
		finisher:    finisher,
	}
}

func (r *RouterGroup) returnObject(ginIRoutes gin.IRoutes) IRoutes {
	switch v := ginIRoutes.(type) {
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
}
