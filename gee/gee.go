package gee

import (
	"log"
	"net/http"
	"strings"
)

//type HandlerFunc func(http.ResponseWriter, *http.Request)
//type Engine struct {
//	//映射表：用来定义路由映射的处理方法
//	router map[string]HandlerFunc
//}
// New is the constructor of gee.Engine
//func New() *Engine {
//	return &Engine{router: make(map[string]HandlerFunc)}
//}
//func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
//	key := method + "-" + pattern
//	engine.router[key] = handler
//}
////会将路由和处理方法注册到映射表 router
//func (engine *Engine) GET(pattern string, handler HandlerFunc) {
//	engine.addRoute("GET", pattern, handler)
//}
//func (engine *Engine) POST(pattern string, handler HandlerFunc) {
//	engine.addRoute("POST", pattern, handler)
//}
//func (engine *Engine) Run(add string) (err error) {
//	return http.ListenAndServe(add, engine)
//}
////ServeHTTP 方法的作用就是，解析请求的路径，查找路由映射表，如果查到，就执行注册的处理方法
//func (engine *Engine) ServeHTTP(w http.ResponseWriter,r *http.Request) {
//	key := r.Method + "-"+r.URL.Path
//	//取出这个函数，然后执行
//	if handler,ok := engine.router[key];ok {
//		handler(w,r)
//	}else {
//		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
//	}
//}
type HandlerFunc  func(*Context)
type Engine struct {
	*RouterGroup
	router *router
	groups[]*RouterGroup // 存放所有分组信息
}
type RouterGroup struct{
	prefix string
	middlewares []HandlerFunc //支持中间件
	parent *RouterGroup
//那么Group对象，还需要有访问Router的能力，
//为了方便，我们可以在Group中，保存一个指针，指向Engine，整个框架的所有资源都是由Engine统一协调的，那么就可以通过Engine间接地访问各种接口了。
	engine *Engine
}
func New()*Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,//TODO
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups,newGroup)
	return newGroup
}
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}
// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}
// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//c := newContext(w, req)
	//engine.router.handle(c)

	var middlewares []HandlerFunc
	for _,group := range engine.groups{
		if strings.HasPrefix(req.URL.Path,group.prefix){
			middlewares = append(middlewares,group.middlewares...)
		}
	}
	c := newContext(w,req)
	c.handlers = middlewares
	engine.router.handle(c)
}
func (group *RouterGroup) Use(middlewares...HandlerFunc)  {
	group.middlewares = append(group.middlewares,middlewares...)
}