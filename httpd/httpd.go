package httpd

import (
	"context"
	"crypto/tls"
	"flag"
	"net/http"
	"time"
)

// 启动https
func StartHttpsServe(ctx context.Context) {
	env := flag.Bool("ENV", true, "Environmental Science")
	flag.Parse()
	r := NewRouter()
	l := &http.Server{
		Addr:         ":9400",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	httpHandMap(r)

	if *env {
		err := l.ListenAndServeTLS("server.crt", "server.key")
		if err != nil {
			panic("https is stoping")
		}
	} else {
		err := l.ListenAndServe()
		if err != nil {
			panic("http is stoping")
		}
	}
}
