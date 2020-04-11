package httpd

import(
	"net/http"
)

func httpHandMap(r *Router) {
	r.Handle("/hello", Hello)
	r.Handle("/world", World)
}

// https 自定义中间件
func makeHandle(h http.) http.Handle {
	return http.Handle(func())
}

// https 挂载函数
func Hello(w http.ResponseWriter, req *http.Request) {
	rsp := NewResponse()
	rsp.Set("this is https web")
	w.Write(rsp.Marshal())
}

func World(w http.ResponseWriter, req *http.Request) {
	rsp := NewResponse()
	rsp.Set("this is https World")
	w.Write(rsp.Marshal())
}
